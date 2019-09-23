# bitcoin-price-reminder

remind you by email when bitcoin price vary 10%

当比特币价格波动超过10%的时候，给你发个QQ邮件

## config file

name: `config.json`

```
{
    "ratio":10.0,
    "email":"x@qq.com",
    "password":"x"
}
```

- `ratio` 价格波动百分比

first time run, it will send u an email to test.

第一次运行的时候，会受到一封测试邮件。

run it every day use crontab.

每天运行一次（使用crontab或者windows计划任务），如果你愿意，也可以每小时运行一次。