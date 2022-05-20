<!--
parent:
  order: false
-->

<div align="center">
  <h1> Cosmos SDK - TR </h1>
</div>

![banner](docs/cosmos-sdk-image.jpg)

<div align="center">
  <a href="https://github.com/cosmos/cosmos-sdk/blob/main/LICENSE">
    <img alt="License: Apache-2.0" src="https://img.shields.io/github/license/cosmos/cosmos-sdk.svg" />
  </a>
  <a href="https://pkg.go.dev/github.com/cosmos/cosmos-sdk?tab=doc">
    <img alt="GoDoc" src="https://godoc.org/github.com/cosmos/cosmos-sdk?status.svg" />
  </a>
  <a href="https://goreportcard.com/report/github.com/cosmos/cosmos-sdk">
    <img alt="Go report card" src="https://goreportcard.com/badge/github.com/cosmos/cosmos-sdk" />
  </a>
  <a href="https://codecov.io/gh/cosmos/cosmos-sdk">
    <img alt="Code Coverage" src="https://codecov.io/gh/cosmos/cosmos-sdk/branch/main/graph/badge.svg" />
  </a>
</div>
<div align="center">
  <a href="https://github.com/cosmos/cosmos-sdk">
    <img alt="Lines Of Code" src="https://tokei.rs/b1/github/cosmos/cosmos-sdk" />
  </a>
  <a href="https://discord.gg/AzefAFd">
    <img alt="Discord" src="https://img.shields.io/discord/669268347736686612.svg" />
  </a>
  <a href="https://sourcegraph.com/github.com/cosmos/cosmos-sdk?badge">
    <img alt="Imported by" src="https://sourcegraph.com/github.com/cosmos/cosmos-sdk/-/badge.svg" />
  </a>
    <img alt="Sims" src="https://github.com/cosmos/cosmos-sdk/workflows/Sims/badge.svg" />
    <img alt="Lint Satus" src="https://github.com/cosmos/cosmos-sdk/workflows/Lint/badge.svg" />
</div>

Cosmos SDK, blok zinciri uygulamaları oluşturmak için bir çerçevedir. Tendermint Core (BFT Consensus) ve Cosmos SDK, Golang programlama dilinde yazılmıştır. Cosmos SDK, Cosmos Hub'ın ilk uygulaması olan Gaia'yı oluşturmak için kullanılır.

UYARI: Cosmos SDK'sı büyük ölçüde stabilize oldu, ancak hala bazı önemli değişiklikler yapıyoruz.

Not: Go 1.18+ gerektirir

## Hemen Başla!

Cosmos SDK'nın üst düzey bir perspektiften nasıl çalıştığını öğrenmek için Cosmos SDK'ya bakın [High-Level Intro](./docs/intro/overview.md).

Hızlı bir şekilde başlamak ve Cosmos SDK'nın üzerine nasıl inşa edileceğini öğrenmek istiyorsanız, adresini ziyaret edin. [Cosmos SDK Tutorials](https://tutorials.cosmos.network). Kendi Cosmos SDK uygulamanızı oluşturmaya başlamak için öğreticinin kaynağını da forklayabilirsiniz.


Daha fazla bilgi için bkz. [Cosmos SDK Documentation](./docs/).

## Dağıtım

bknz. [CONTRIBUTING.md](./CONTRIBUTING.md) nasıl katkıda bulunacağınız ve katılacağınız ayrıntılar için bknz. [dev calls](./CONTRIBUTING.md#teams-dev-calls).

Güncellemeleri takip etmek veya en son tasarım hakkında daha fazla bilgi edinmek istiyorsanız, bize katılın. [Discord](https://discord.com/invite/cosmosnetwork).

## Araçlar and Frameworks

Cosmos ekosistemi çok geniştir. Burada sadece dikkate değer birkaç söz edeceğiz.

+ [Tools](https://v1.cosmos.network/tools): dikkate değer Frameworks ve modüller.
+ [CosmJS](https://github.com/cosmos/cosmjs): JavaScript tabanlı istemci çözümlerini güçlendirmek için İsviçre Çakısı.

### Cosmos Hub Mainnet

Cosmos Hub uygulaması 'gaia' kendi kendine taşındı [cosmos/gaia repository](https://github.com/cosmos/gaia). Cosmos Hub ana ağına ve daha fazlasına katılmak için oraya gidin.

### Inter-Blockchain Communication (IBC)

Cosmos SDK için IBC modülü kendi modülüne taşındı [cosmos/ibc-go repository](https://github.com/cosmos/ibc-go). IBC modülünü oluşturmak ve entegre etmek için oraya gidin.
### Ignite CLI

Ignite CLI, bağımsız ve güvenli bir blok zincirinde herhangi bir kripto uygulamasını oluşturmak, başlatmak ve sürdürmek için hepsi bir arada platformdur. Yeni bir uygulama veya yeni bir modül oluşturuyorsanız, şunu kullanın:[Ignite CLI](https://github.com/ignite-hq/cli) başlamak ve gelişmeyi hızlandırmak için.

## Disambiguation

Bu Cosmos SDK projesi ile ilgili değil [React-Cosmos](https://github.com/react-cosmos/react-cosmos) proje (henüz). Bu Github organizasyon adı için Evan Coury ve Ovidiu (@skidding)'e çok teşekkürler. Anlaşmamıza göre, bu anlam ayrımı bildirimi burada kalacak.
