# yURL: Universal Links / AASA File Validator

[![CircleCI Build Status](https://circleci.com/gh/chayev/yurl.svg?style=shield)](https://circleci.com/gh/chayev/yurl) [![GitHub License](https://img.shields.io/badge/license-MIT-blue.svg)](https://raw.githubusercontent.com/chayev/yurl/main/LICENSE)

yURL is a CLI (Command-Line Interface) and [webapp](https://yurl.chayev.com/) that allows you to validate whether a URL is properly configured for Universal Links and App Links. This allows you to check if the apple-app-site-association (AASA) and assetlinks.json files exist and are in the proper format as defined by [Apple](https://developer.apple.com/documentation/safariservices/supporting_associated_domains) and [Google](https://developers.google.com/digital-asset-links/v1/getting-started).

## macOS Install Instructions

### Install with Brew (recommended)

Install yURL with [Brew](https://brew.sh/):

```
brew install chayev/tap/yurl
```

### Install using cURL

Run the below command:

```
curl -sSL "https://github.com/tkhoa2711/yurl/releases/download/v0.7.3/yurl-v0.7.3-macos-amd64.tar.gz" | sudo tar -xz -C /usr/local/bin yurl
```

Note: You will be prompted to enter your password because of the `sudo` command. `0.7.3` may need to be replaced with your desired version.

## Linux Install Instructions

### Install using cURL (recommended)

Run the below command:

```
curl -sSL "https://github.com/tkhoa2711/yurl/releases/download/v0.7.3/yurl-v0.7.3-linux-amd64.tar.gz" | sudo tar -xz -C /usr/local/bin yurl
```

Note: You will be prompted to enter your password because of the `sudo` command. `0.7.3` may need to be replaced with your desired version.

### Install with Snap (deprecated)

Install yURL with [Snap](https://snapcraft.io/):

```
sudo snap install yurl
```

## Windows Install Instructions

You could download the executable from the [releases](https://github.com/tkhoa2711/yurl/releases) page. More instructions coming soon!

We are planning on supporting [chocolatey](chocolatey.org) package manager as well.

## Usage and Example

Run `yurl help` for information on how to use yURL.

Example:

To validate the Apple App Site Association (AASA) file run the following:

```bash
yurl aasa validate suadeo.onelink.me
```

```bash
yurl aasa check ./path/to/local/aasa/file
```

To validate the Android assetlinks.json file run the following:

```bash
yurl assetlink validate suadeo.onelink.me
```

## Contributing

Contributions to yURL of any kind are welcome! Feel free to open [PRs](https://github.com/tkhoa2711/yurl/pulls) or an [issue](https://github.com/tkhoa2711/yurl/pulls).

### Asking Support Questions

Feel free to open an issue if you have a question.

### Reporting Issues

If you believe you have found a defect in yURL or its documentation, create an issue to report the problem.
When reporting the issue, please provide the version of yURL in use (`yurl version`).

## License

This repository is licensed under the MIT license.
The license can be found [here](./LICENSE).
