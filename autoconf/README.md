# autoconf

This build step provides a container with the build-essentials and autoconf
packages installed. It is intended for running automake targets. Git is included
as well since autoconf scripts tend to read git status for baking version info
into programs.

It is capable of running simple GCC builds, but the primary use is in conjuction
with other containers as part of a Google Container Builder configuration.

The entrypoint for this container is autoconf.
