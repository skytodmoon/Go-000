基于 errgroup 实现一个 http server 的启动和关闭 ，以及 linux signal 信号的注册和处理，要保证能够一个退出，全部注销退出。

// Package errgroup provides synchronization, error propagation, and Context
// errgroup 包为一组子任务的 goroutine 提供了 goroutine 同步,错误取消功能.
//
//errgroup 包含三种常用方式
//
//1、直接使用 此时不会因为一个任务失败导致所有任务被 cancel:
//		g := &errgroup.Group{}
//		g.Go(func(ctx context.Context) {
//			// NOTE: 此时 ctx 为 context.Background()
//			// do something
//		})
//
//2、WithContext 使用 WithContext 时不会因为一个任务失败导致所有任务被 cancel:
//		g := errgroup.WithContext(ctx)
//		g.Go(func(ctx context.Context) {
//			// NOTE: 此时 ctx 为 errgroup.WithContext 传递的 ctx
//			// do something
//		})
//
//3、WithCancel 使用 WithCancel 时如果有一个人任务失败会导致所有*未进行或进行中*的任务被 cancel:
//		g := errgroup.WithCancel(ctx)
//		g.Go(func(ctx context.Context) {
//			// NOTE: 此时 ctx 是从 errgroup.WithContext 传递的 ctx 派生出的 ctx
//			// do something
//		})
//

Errorgroup 资料
https://github.com/golang/sync/blob/09787c993a3a/errgroup/errgroup.go
https://github.com/go-kratos/kratos/blob/76da31effb5ece597cff22e970816a5ddd7a7659/pkg/sync/errgroup/errgroup.go
https://zhuanlan.zhihu.com/p/64983626