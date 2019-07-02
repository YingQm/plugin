// Copyright Fuzamei Corp. 2018 All Rights Reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package executor

/*
该合约主要是配合Dpos共识，完成（1）候选节点的注册、去注册、投票及查询管理。（2）Dpos共识运行过程中，得票数TopN（N为约定的受托节点数量）受托节点的VRF相关信息的分阶段发布管理。

（1）系统初始运行时，会有默认的几个受托节点进行共识运行。
（2）系统运行后，可以重新选举受托节点，各个候选节点需要抵押10000个币（暂未实现），注册成为候选节点。
（3）候选节点可以在社区宣传，让大家为自己投票。
（4）用户可以为自己支持的候选节点投票。投票后，资金会冻结，3天以后才可以撤销投票。
（5）系统运行过程中，每到固定区块高度时（比如10万个区块），重新获取当前投票数据，并确定下一个时间段的受托节点。
（6）受托节点进行共识出块，每个cycle（一个cycle中，各个受托节点轮番出块，直到都轮一遍）分两个阶段进行VRF信息发布：
     第一个阶段，各个受托节点发布自己的VRF的M信息
     第二个阶段，各个受托节点发布自己的VRF的R、P信息
     上述VRF的M、R、P可以验证。
（7）新的cycle中，使用上述VRF信息进行受托节点的出块顺序的重新洗牌，按洗牌结果决定各受托节点出块的顺序
*/
