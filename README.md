## 实现分布式，水平扩展的游戏聊天服务器 scalable chat im

## archi
- [ ] gateway
    * 负责对client的所有请求
    * 尽量逻辑简单，单gateway能够承载足够多的用户，为了减少整体gateway的数量
- [ ] pusher
    * 消息推送，消息分发，无状态
- [ ] logic
    * 一致性验证，关系验证，消息验证，消息存储等

### feature
- [ ] 消息以channel为单位
    * 队伍-工会-世界为注册channel
    * 点聊为分发

### message-flow
- [ ] 仅允许一端登录
  * |gateway-|auth-|gateway-|logic-lock-kickother-|gateway lock采用乐观锁   

### todolist
- [ ] 
- [ ] 登录
    -[ ] gateway
    -[ ] logic 
    -[ ] auth