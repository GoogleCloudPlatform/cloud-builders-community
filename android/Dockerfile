FROM openjdk:8-slim
ENV DEBIAN_FRONTEND noninteractive
ENV LANG='en_US.UTF-8' LANGUAGE='en_US:en' LC_ALL='en_US.UTF-8'

ARG ANDROID_SDK_TOOLS_URL="https://dl.google.com/android/repository/sdk-tools-linux-3859397.zip"
ARG ANDROID_SDK_LICENSE
ARG ANDROID_BUILD_CACHE
ARG GRADLE_DISTRIBUTION="https://services.gradle.org/distributions/gradle-4.4-bin.zip"

ENV GRADLE_BUILD_CACHE=${ANDROID_BUILD_CACHE}
ENV DOCKER_ANDROID_DISPLAY_NAME mobileci-docker
ENV GRADLE_TARGETS="assembleDebug"
ENV GRADLE_HOME=/usr/local/gradle-4.4

ENV ANDROID_HOME /android-sdk
ENV PATH ${ANDROID_HOME}/tools:$ANDROID_HOME/platform-tools:$ANDROID_HOME/tools/bin:/bin:$PATH
#Install underlying development packages and languages
ADD packages.txt . 
COPY gradle-build /bin

    # Prepare to install packages
RUN apt-get update && \
    apt-get install -y locales && \
    sed -i 's/^# *\(en_US.UTF-8\)/\1/' /etc/locale.gen && \
    ln -s /etc/locale.alias /usr/share/locale/locale.alias && \ 
    locale-gen && \
    dpkg-reconfigure locales && \
    apt-get install -y apt-utils && \
    dpkg --add-architecture i386 && \
    # install packages
    apt-get install -qq -y \
    software-properties-common \
    curl \
    wget \
    build-essential \
    zip \
    unzip \
    --no-install-recommends && \
    # Install gradle
    wget --quiet $GRADLE_DISTRIBUTION && \
    unzip -qq  gradle-4.4-bin.zip -d /usr/local/ && \
    ln -s /usr/local/gradle-4.4/bin/gradle /usr/local/bin/gradle && \
    # Install android sdk
    mkdir -p /root/.android && touch /root/.android/repositories.cfg && \
    wget --quiet --output-document=android-sdk.zip ${ANDROID_SDK_TOOLS_URL} && \
    unzip -qq android-sdk.zip -d $ANDROID_HOME && \
    mkdir $ANDROID_HOME/licenses && \
    echo ${ANDROID_SDK_LICENSE} >> $ANDROID_HOME/licenses/android-sdk-license && \
    echo "sdk.dir=${ANDROID_HOME}" > android_local.properties && \
    sdkmanager --package_file=packages.txt && \
    chmod -R +x bin && \ 
    rm $ANDROID_HOME/licenses/android-sdk-license 
CMD ["sh", "-c", "build-debug", "${GRADLE_TARGETS}"]
