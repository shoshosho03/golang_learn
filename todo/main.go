package main

// 必要なパッケージのインポート
import (
	"encoding/json"
	"fmt"
	"os"
)

// Todo構造体の定義
type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

// JSONファイルの名前を定数として定義
const fileName = "todos.json"

// メイン関数
func main() {
	// コマンドライン引数のチェック
	if len(os.Args) < 2 {
		// コマンドの使用方法を表示
		fmt.Println("Usage: go run main.go [add|list] [todo text]")
		return
	}
	// コマンドを取得
	command := os.Args[1]

	// コマンドに応じた処理を実行
	switch command {
	// "add"コマンドの場合
	case "add":
		// 引数が不足している場合のエラーメッセージ
		if len(os.Args) < 3 {
			fmt.Println("Please provide a todo text.")
			return
		}
		// 追加するTodoのテキストを取得
		text := os.Args[2]
		addTodo(text)
	// "list"コマンドの場合
	case "list":
		listTodos()
	// その他のコマンドの場合のエラーメッセージ
	default:
		fmt.Println("Unknown command. Use 'add' or 'list'.")
	}
}

// Todoを追加する関数
func addTodo(text string) {
	// 既存のTodoをロード
	todos := loadTodos()
	// 新しいTodoのIDを生成
	id := len(todos) + 1
	// 新しいTodoを作成してリストに追加
	todo := Todo{ID: id, Text: text}
	// Todoを保存
	todos = append(todos, todo)
	saveTodos(todos)
	// 追加したTodoの情報を表示
	fmt.Printf("Added todo: %s\n", text)
}

// Todoのリストを表示する関数
func listTodos() {
	// 既存のTodoをロード
	todos := loadTodos()
	// Todoが存在しない場合のメッセージ
	if len(todos) == 0 {
		fmt.Println("No todos found.")
		return
	}
	// Todoのリストを表示
	for _, todo := range todos {
		fmt.Printf("%d: %s\n", todo.ID, todo.Text)
	}
}

// TodoをJSONファイルからロードする関数
func loadTodos() []Todo {
	// Todoのリストを初期化
	var todos []Todo
	// JSONファイルを開く
	file, err := os.Open(fileName)
	// ファイルが存在しない場合は空のリストを返す
	if err != nil {
		if os.IsNotExist(err) {
			return todos
		}
		// その他のエラーの場合はエラーメッセージを表示して空のリストを返す
		fmt.Println("Error opening file:", err)
		return todos
	}
	// ファイルを閉じるためのdefer文
	defer file.Close()
	// JSONデコードを行う
	err = json.NewDecoder(file).Decode(&todos)
	// デコードエラーが発生した場合のエラーメッセージ
	if err != nil {
		fmt.Println("Error decoding JSON:", err)
	}
	return todos
}

// TodoをJSONファイルに保存する関数
func saveTodos(todos []Todo) {
	// JSONファイルを作成または上書きする
	file, err := os.Create(fileName)
	// ファイル作成エラーが発生した場合のエラーメッセージ
	if err != nil {
		// エラーメッセージを表示して関数を終了
		fmt.Println("Error creating file:", err)
		return
	}
	// ファイルを閉じるためのdefer文
	defer file.Close()
	// JSONエンコードを行う
	err = json.NewEncoder(file).Encode(todos)
	// エンコードエラーが発生した場合のエラーメッセージ
	if err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}
