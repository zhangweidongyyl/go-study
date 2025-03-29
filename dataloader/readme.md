batcher
功能
该代码实现了一个Gin框架下的通用批处理中间件，核心功能是将多个数据请求合并成批量操作，减少对后端服务的高频调用。适用于解决N+1查询问题，例如在GraphQL或REST API中需要批量获取资源时，自动合并请求。

设计
1. 核心结构：
   ○ batchFactory：全局工厂，通过sync.Map存储所有批处理器（batcher），每个batcher由唯一key标识。
   ○ batcher：实际执行批处理的单元，管理请求队列、触发条件和结果分发。
   ○ batchRequest：单个请求的封装，包含键（key）和结果通道（channel）。
2. 关键流程：
   ○ 注册中间件：通过Register()在每个Gin请求的上下文中注入batchFactory。
   ○ 获取批处理器：通过batch()函数按key获取或创建batcher，支持配置容量（cap）、等待时间（wait）和追踪器（tracer）。
   ○ 收集请求：请求通过collect()加入队列，触发条件为队列满或超时。
   ○ 执行批量操作：调用用户定义的BatchFunc获取数据，处理错误或Panic，并通过通道分发结果。
3. 并发控制：
   ○ 使用sync.Mutex保护requests队列的并发写入。
   ○ 通过sync.Map实现线程安全的batcher存储。
   ○ 每个批处理通过独立的Goroutine执行，避免阻塞主流程。
4. 扩展性：
   ○ Tracer接口：允许用户跟踪批处理生命周期（开始、结束）。
   ○ 泛型支持：通过K（键类型）和V（值类型）支持不同数据类型的批处理。