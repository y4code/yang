# yang
羊了个羊脚本

在https://github.com/zc2638/ylgy 的基础上加了高并发，但注意次数不要太高，一般操作系统 socket 上限并不高，报错的话就把 times 调小一点，默认 100

## 下载
 - windows 下载 https://github.com/y4code/yang/raw/main/yang.exe

 - mac 下载 https://github.com/y4code/yang/raw/main/yang

## 用法
直接点开文件根据提示输入token回车，或者命令行如下
```
./yang -times 100
token 为空, 请继续输入token, 或者直接使用 -token 参数
（粘贴token回车）
```

或者

```
./yang -times 100 -token 你的token

```
