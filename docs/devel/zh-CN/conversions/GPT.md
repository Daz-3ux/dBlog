## 调用 OPENAI 思路
- 使用 LangChain Go 构建一个简单的调用 OPENAI 的服务
  - 什么是 LangChain? 
    - ⚡ Building applications with LLMs through composability ⚡
- 使用[模型](https://platform.openai.com/docs/models/overview): 可自行设置
- API_TOKEN: 自行设置
- 示例:
```sql
export OPENAI_API_KEY='sk-xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx'
export OPENAI_MODEL='GPT-3.5-Turbo'
```

## 作用
- 总结博客内容以及生成简易提纲
  - 方便用户快速了解博客内容
  - 在快餐化时代,人们的阅读习惯也在发生变化,人们更倾向于快速了解内容,而不是深入阅读
  - sad, but true and useful

## 核心代码
```go
func createAIComment(postContent string) (string, error) {
	llm, err := openai.NewChat(openai.WithModel(os.Getenv("OPENAI_MODEL")), openai.WithToken(os.Getenv("OPENAI_API_KEY")))
	if err != nil {
		return "", err
	}

	chatMsg := []schema.ChatMessage{
		schema.HumanChatMessage{
			Content: "总结以下内容,生成一份简短的提纲以及内容摘要,不要有多余输出" + postContent,
		},
	}

	aiMsg, err := llm.Call(context.Background(), chatMsg)
	if err != nil {
		log.C(context.Background()).Errorw("failed to call AI", "error", err)
		return "", err
	}

	return aiMsg.GetContent(), nil
}
```