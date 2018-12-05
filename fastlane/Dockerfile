FROM ruby:2.5.3-alpine

RUN apk add --update --no-cache ruby-dev build-base libxml2-dev libxslt-dev
RUN gem i fastlane --no-ri --no-rdoc

ENV LC_ALL=en_US.UTF-8 LANG=en_US.UTF-8

ENTRYPOINT ["fastlane"]
CMD []
