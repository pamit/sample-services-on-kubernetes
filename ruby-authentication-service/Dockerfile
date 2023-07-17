FROM ruby:3.1.2

RUN apt-get update -qq && apt-get install -y build-essential

WORKDIR /auth_service
COPY . /auth_service
RUN bundle install

EXPOSE 4567

CMD ["bundle", "exec", "rackup", "--host", "0.0.0.0", "-p", "4567"]
