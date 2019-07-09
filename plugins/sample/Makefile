PLUGINNAME=sample
PLUGIN=$(PLUGINNAME).px

all: $(PLUGIN)

$(PLUGIN):
	go build -tags=plugin -o $(PLUGIN) -buildmode=plugin

clean:
	rm $(PLUGIN)

.PHONY: all clean

