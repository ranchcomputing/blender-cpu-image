# Blender docker images for CPU rendering

[![GitHub tag (with filter)](https://img.shields.io/github/v/tag/ranchcomputing/blender-cpu-image?logo=docker&label=latest)](https://github.com/ranchcomputing/blender-cpu-image/pkgs/container/blender-cpu-image)


For instructions go to the [tagged image versions list](https://github.com/ranchcomputing/blender-cpu-image/pkgs/container/blender-cpu-image).


## Release process

New blender versions are checked daily from the [github mirror](https://github.com/blender/blender/tags).

In case of a failure, you can create a `pre/vX.X.X` branch and debug the action.

When ready, [create](https://github.com/ranchcomputing/blender-cpu-image/releases/new) a `vX.X.X` tag or manually trigger the [Auto-update workflow](https://github.com/ranchcomputing/blender-cpu-image/actions/workflows/cron.yml).

---

The Dockerfile was inspired by https://github.com/nytimes/rd-blender-docker

Blender download: https://www.blender.org/download/

Blender releases: https://download.blender.org/release/
