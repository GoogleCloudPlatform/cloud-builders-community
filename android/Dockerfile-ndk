ARG project_id
ARG android_version
FROM gcr.io/${project_id}/android:${android_version}

ARG ndk_version=android-ndk-r17b
ARG android_ndk_home=/opt/android/${ndk_version}

# If we're using the NDK we likely want cmake
RUN sdkmanager "cmake;3.6.4111459"

# Install the NDK
RUN curl --silent --show-error --location --fail --retry 3 --output /tmp/${ndk_version}.zip https://dl.google.com/android/repository/${ndk_version}-linux-x86_64.zip && \
    sudo unzip -q /tmp/${ndk_version}.zip -d /opt/android && \
    rm /tmp/${ndk_version}.zip

ENV ANDROID_NDK_HOME ${android_ndk_home}
