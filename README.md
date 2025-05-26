说明
==============
只能用来支付THB泰铢,且每次至少要300泰铢才可以.
(所以参数中不需要指定currency)

Deposit
==============
1. 生成一个支付qrcode二维码
2. 用户扫描支付
3. 回调通知merchant (回调接口是merchant商户申请时即指定的)

Withdraw
==============
用户发起即可, 银行收到后会直接执行

Comment
===============
both support deposit && withdrawl


鉴权
==============
1. 请求时在header里放了partner_code 和 authorization . 通过这个进行鉴权 (并没有对请求body做签名)
2. 充值回调: 是我们自己做了一个md5签名放到了ref4字段里, 通过这个来预防假的callback


回调地址
==============
是提前让12pay设置好的.（无法api中动态修改）
TODO 这里得有个后台??