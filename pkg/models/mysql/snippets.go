package mysql

import (
	"database/sql"
	"errors"

	"github.com/maxfeizi04-cloude/snippetbox/pkg/models"
)

// SnippetModel ,封装 sql.DB 连接池
type SnippetModel struct {
	DB *sql.DB
}

// Insert 向数据库中插入一条新 snippet
func (m *SnippetModel) Insert(title, content, expires string) (int, error) {
	// 编写要执行的 SQL 语句。为便于阅读，我将其分成了两行
	//(因此用反引号括起来，而不是普通的双引号)
	stmt := `INSERT INTO snippets (title, content, created, expires)
	VALUES (?,?,UTC_TIMESTAMP(),DATE_ADD(UTC_TIMESTAMP(), INTERVAL ? DAY ))`

	// 使用内嵌连接池的 Exec() 方法执行该语句
	// 第一个参数是 SQL 语句，后面依次是占位符参数对应的 title、content 和 expires 值
	// 该方法返回一个 sql.Result 对象，其中包含语句执行情况的基本信息
	result, err := m.DB.Exec(stmt, title, content, expires)
	if err != nil {
		return 0, err
	}
	// 使用 result 对象的 LastInsertId() 方法获取 snippets 表中新插入记录的 ID
	id, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	// 返回的 ID 类型为 int64，因此返回前将其转换为 int 类型
	return int(id), nil
}

// Get 根据 id 返回对应的 snippet
func (m *SnippetModel) Get(id int) (*models.Snippet, error) {

	// 编写要执行的 SQL 语句
	// 同样,为便于阅读,我将其分成了两行
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() AND id = ?`

	// 使用连接池的 QueryRow() 方法执行 SQL 语句,将不可信的 id 变量作为占位符参数的值传入
	// 该方法返回一个指向 sql.Row 对象的指针,该对象包含数据库返回的结果
	row := m.DB.QueryRow(stmt, id)

	// 初始化一个指向零值 Snippet 结构体的指针
	s := &models.Snippet{}

	// 使用 row.Scan() 将 sql.Row 中各字段的值复制到 Snippet 结构体对应的字段中
	// 注意,row.Scan 的参数是指向目标位置的指针
	// 且参数数量必须与语句返回的列数完全一致
	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		// 如果查询没有返回行，row.Scan() 会返回 sql.ErrNoRows 错误
		// 我们使用 errors.Is() 函数专门检查该错误，并返回自定义的 models.ErrNoRecord 错误
		if errors.Is(err, sql.ErrNoRows) {
			return nil, models.ErrNoRecord
		}
		return nil, err
	}
	// 一切正常则返回 Snippet 对象。
	return s, nil
}

// Latest 返回最近创建的 10 条 snippet
func (m *SnippetModel) Latest() ([]*models.Snippet, error) {
	// 编写要执行的 SQL 语句
	stmt := `SELECT id, title, content, created, expires FROM snippets
	WHERE expires > UTC_TIMESTAMP() ORDER BY created DESC LIMIT 10`
	// 使用连接池的 Query() 方法执行 SQL 语句
	// 该方法返回一个包含查询结果的 sql.Rows 结果集
	rows, err := m.DB.Query(stmt)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// 初始化一个空切片，用于存放 models.Snippet 对象
	snippets := []*models.Snippet{}

	// 使用 rows.Next() 遍历结果集中的行
	// 这会准备第一行（及随后的每一行），供 rows.Scan() 方法处理
	// 若遍历完成所有行，结果集会自动关闭并释放底层数据库连接
	for rows.Next() {
		// 创建一个指向零值 Snippet 结构体的指针
		s := &models.Snippet{}

		// 使用 rows.Scan() 将行中各字段的值复制到新创建的 Snippet 对象中
		// 同样，row.Scan() 的参数必须是指向目标位置的指针
		// 且参数数量必须与语句返回的列数完全一致
		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		// 将其追加到 snippets 切片中。
		snippets = append(snippets, s)
	}

	// rows.Next() 循环完成后，调用 rows.Err() 获取迭代过程中遇到的任何错误
	// 调用此方法很重要——不要假设已成功完成对整个结果集的迭代
	if err = rows.Err(); err != nil {
		return nil, err
	}
	// 一切正常则返回 Snippets 切片
	return snippets, nil
}
