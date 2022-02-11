<div align="center">
  <a href="http://www.amitmerchant.com/electron-markdownify"><img src="https://www.clipartmax.com/png/full/33-336258_drawing-clipart-transparent-polaroid-clipart.png" alt="Markdownify" width="200"></a>
  <h1>iReaver</h1>
</div>

<div align="center">
    <p>
        <strong>
            Mac Photos backup export tool. This tool will export all photos in your Photos application<br> to a directory
            of your choice with 3 naming options to choose from.
        </strong>
    </p>
</div>

<div align="center">
    <img src="https://github.com/Celtech/iReaver/actions/workflows/release.yml/badge.svg"
     alt="Pipeline Status">
    <img src="https://img.shields.io/badge/apache%20version-2.4-orange?&logo=apache">
    <img src="https://img.shields.io/badge/binary%20size-1.4%20MB-blue?&logo=go">
</div>

<div align="center">
  <a href="#installation">Installation</a> •
  <a href="#quick-start">Quick Start</a> •
  <a href="#flags">Flags</a>
</div>

## Installation

### Homebrew

```shell
$ brew tap Celtech/iReaver
$ brew install iReaver
```

### Go Install

```shell
$ go install github.com/Celtech/iReaver@latest
```

### Manual

You can download the pre-compiled binary for you specific OS from the [releases page](https://github.com/Celtech/iReaver/releases). 
You will need to copy the and extract the binary, then move it to your local bin folder. Please, refer to the example below:
```shell
curl https://github.com/Celtech/iReaver/releases/download/${VERSION}/${PACKAGE_NAME} -o ${PACKAGE_NAME}
sudo tar -xvf ${PACKAGE_NAME} -C /usr/local/bin/
sudo chmod +x /usr/local/bin/iReaver
```

## Quick Start
Before we can use iReaver, we first need to configure a setting inside the photos application. Open the Photos application 
on your Mac. Once open, click `Photos -> Prefernces...` in the menu bar. Once your preferences window opens, head to the 
iCloud tab, and ensure that `Download Originals to this Mac` is selected.

<img width="700" alt="Screen Shot 2022-02-10 at 11 12 44 PM" src="https://user-images.githubusercontent.com/7949140/153540857-3148ffd7-fe84-4ad8-8cd7-4e1652983f99.png">

Once you have this done, wait for Photos to finish downloading all the files. After that let's run the following command:

```shell
$ ireaver -o "~/photos-backup"
```

This command will copy all the files from your Photos app to a folder in your home directory using the original file names. 
For a complete list of all possible flags, see the flags section below.

## Flags
```shell
  -h, --help            help for ireaver
  -n, --name string     Renames output images (options: "sequential", "original", "guid") (default "original")
  -o, --output string   Directory path to the folder you want to output images to
  -p, --path string     Path to your local iCloud photos library (default "$HOME/Pictures/Photos Library.photoslibrary/")
```

### name
The name flag is used to modify the copied file name pattern.

| Option     | Description                                                                                                               |
|------------|---------------------------------------------------------------------------------------------------------------------------|
| sequential | Names the files in sequential order based on their index, this will guarantee the newest files are always the last index. |
| original   | Names the files using their original names as displayed in the photos app, this is the default. I.e `img_1234.heic`       |
| guid       | Names the files using the asset GUID, note images will be in a random order if used.                                      |

### output
The output directory of the copied files. This flag supports both `~` and `$HOME` which represent the users home directory.
It does not matter if the file path ends with a `/` or not.

### path
The directory where your Photos application lives. On most Macs this is `$HOME/Pictures/Photos Library.photoslibrary/`,
and this is the default path.