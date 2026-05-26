#!/usr/bin/env nu

const root = path self .
let release = http get https://api.github.com/repos/rakaly/librakaly/releases/latest

def download [name_suffix: string files: list<string>] {
  let asset = $release.assets | where name ends-with $name_suffix | first
  let url = $asset.browser_download_url
  let dir_name = $asset.name | path parse --extension "tar.gz" | get stem
  let archive_file = mktemp --dry
  let build_dir = mktemp --directory
  try {
    http get $url | save $archive_file
    bsdtar -xf $archive_file -C $build_dir --strip-components 1 $dir_name
    mv ...($files | each {|file| $"($build_dir)/($file)" }) $root
  } catch {|err|
    print -n $err.rendered
  }
  rm -rf $archive_file $build_dir
}

download "-linux.tar.gz" [librakaly.so]
download "-macos.tar.gz" [librakaly.dylib]
download "-win-msvc.tar.gz" [rakaly.dll rakaly.h]

$'($release.tag_name)
https://github.com/rakaly/librakaly/releases/latest
' | save -f $"($root)/VERSION"
