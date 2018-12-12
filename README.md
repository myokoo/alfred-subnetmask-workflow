# alfred-subnetmask-workflow

Alfred workflow to calculate subnet mask.

#### subnetmask xxx.xxx.xxx.xxxx/xx
![subnetmask](https://user-images.githubusercontent.com/1995330/49699178-69888b80-fc11-11e8-9cf5-56b16b9263c9.gif)


#### subnetmask xxx.xxx.xxx.xxx
![broadcast](https://user-images.githubusercontent.com/1995330/49859147-26f5c780-fe3a-11e8-94b0-e2234e967d0a.gif)

#### subnetmask /xx
![prefix](https://user-images.githubusercontent.com/1995330/49699181-70af9980-fc11-11e8-931d-1f8d5a545e9c.gif)



### Setup

#### go version >= 1.11.1

```bash
git clone git@github.com:myokoo/alfred-subnetmask-workflow.git
go build

sh setup.sh
```

`setup.sh` is created link to alfred workflow v3 directory.

### Licensing
This library is released under the [MIT license][license].

The icon is <a href="http://creativecommons.org/licenses/by/4.0/" rel="license"><img style="border-width: 0; font-size: 15px; vertical-align: middle;" src="https://i.creativecommons.org/l/by/4.0/80x15.png" alt="クリエイティブ・コモンズ・ライセンス" /></a> [SAKURA Internet Inc.](https://knowledge.sakura.ad.jp/4724/)

[license]: ./LICENSE
