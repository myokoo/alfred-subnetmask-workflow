# alfred-subnetmask-workflow

Alfred workflow to calculate subnet mask.

#### subnetmask xxx.xxx.xxx.xxxx/xx
![subnetmask](https://user-images.githubusercontent.com/1995330/48602926-1deb1500-e9b8-11e8-9e17-fd12e8f3aae4.gif)


#### subnetmask xxx.xxx.xxx.xxx
![broadcast](https://user-images.githubusercontent.com/1995330/48602931-22173280-e9b8-11e8-9d5e-132837ffc902.gif)

#### subnetmask /xx
![prefix](https://user-images.githubusercontent.com/1995330/48602937-25aab980-e9b8-11e8-9359-cb7e917ac2c0.gif)



### Setup

#### go version >= 1.10.3

```bash
git clone git@github.com:myokoo/alfred-subnetmask-workflow.git
go get golang.org/x/vgo
vgo build

sh setup.sh
```

`setup.sh` is created link to alfred workflow v3 directory.

### Licensing
This library is released under the [MIT license][license].

The icon is <a href="http://creativecommons.org/licenses/by/4.0/" rel="license"><img style="border-width: 0; font-size: 15px; vertical-align: middle;" src="https://i.creativecommons.org/l/by/4.0/80x15.png" alt="クリエイティブ・コモンズ・ライセンス" /></a> [SAKURA Internet Inc.](https://knowledge.sakura.ad.jp/4724/)

[license]: ./LICENSE
