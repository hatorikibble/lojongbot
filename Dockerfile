FROM iron/base

# export CGO_ENABLED=0 



ADD lojongbot /
ADD config.json /
ADD slogans.txt /

ENTRYPOINT ["./lojongbot"]