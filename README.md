# pandora-wechat
通过 pandora ，无需翻墙即可将ChatGPT接入微信

通过 pandora 访问ChatGPT，呼吸更加顺畅了~

本项目通过 https://github.com/eatmoreapple/openwechat 开源项目提供微信聊天回复支持。暂不支持语音处理。

本项目通过 https://github.com/pengzhile/pandora 开源项目提供ChatGPT访问支持。

本项目由GitHub Copilot提供代码编写支持。
# 接入方式：
首先，前往 https://github.com/pengzhile/pandora 项目，按照说明，将pandora部署到服务器上（部署在本地PC也可以，推荐docker部署方式，非常方便）[请注意：目前所支持的pandora版本为1.0.10版本，高版本pandora可能出现兼容性问题。]

然后，从本项目的发行版(Release)中下载对应系统的可执行程序，并将其重命名为 pandora-wechat。

无需任何依赖环境，直接运行./pandora-wechat即可

之后，该程序会在控制台输出网页链接，将该链接复制到浏览器，通过微信扫描，确认登录即可。
# 高级：
Pro 用户可以自己修改本开源项目中的任何代码，并通过安装 go 语言环境，自行构建部署。
