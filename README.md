# GoMemory

该项目目前仅用于向pinecone向量库中批量导入向量。仅包含此功能。

## 快速开始

- 填写config目录下的config.cfg。要求有OpenAI的api-key，piencone的api-key与index的url
- 将想要批量导入的预处理语句写入store下的text文件。以A、B作为类型分隔符，空格作为句内元素分隔符。A代表memory，B代表knowledge。
- 运行程序，开始批量导入，并将导入的结果以id:text写入`store/log`里以运行开始时间生成的txt文件当中。