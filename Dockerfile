FROM ubuntu:22.04

RUN apt-get update && apt-get install -y \
    wget \
    # to extract tar.xz
	xz-utils \
    # from https://github.com/blender/blender/blob/v3.3.0/release/freedesktop/snap/blender-snapcraft-template.yaml
    libxcb1 \
    libxext6 \
    libx11-6 \
    libxi6 \
    libxfixes3 \
    libxrender1 \
    libxxf86vm1 \
    # libGL
    libgl1-mesa-glx \
    # libxkbcommon.so.0 needed since 3.4.0 apparently
    libxkbcommon0 \
    # libSM.so.6 needed since 3.5.1 apparently
    libsm6

# https://ftp.halifax.rwth-aachen.de/blender/release
ARG BLENDER_MIRROR=https://mirror.clarkson.edu/blender/release
ARG BLENDER_VERSION=3.3.0
ARG BLENDER_MAJOR

RUN export BLENDER_MAJOR="${BLENDER_MAJOR:-${BLENDER_VERSION%.*}}" && echo ${BLENDER_MIRROR}/Blender${BLENDER_MAJOR}/blender-${BLENDER_VERSION}-linux-x64.tar.xz && \
    wget --no-verbose --show-progress --progress=dot:giga \
    ${BLENDER_MIRROR}/Blender${BLENDER_MAJOR}/blender-${BLENDER_VERSION}-linux-x64.tar.xz \
	&& tar -xvf blender-${BLENDER_VERSION}-linux-x64.tar.xz --strip-components=1 -C /bin \
	&& rm -rf blender-${BLENDER_VERSION}-linux-x64.tar.xz \
	&& rm -rf blender-${BLENDER_VERSION}-linux-x64

RUN blender --version

# make symlink to /bin/python
RUN echo 'import os, sys; os.symlink(sys.executable, "/bin/python")' | blender --background --python-console
RUN python --version
