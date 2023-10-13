# Blender docker images for CPU rendering

Container image versions: https://github.com/oliverpool/blender-cpu-image/pkgs/container/blender-cpu-image

Blender versions: https://download.blender.org/release/ or https://github.com/blender/blender/tags

## Release process

To ensure that the build process works, you can push a `pre/v3.6.0` branch and check the action.

When ready, [create](https://github.com/oliverpool/blender-cpu-image/releases/new) a `vX.X.X` tag.

---

The Dockerfile was inspired by https://github.com/nytimes/rd-blender-docker