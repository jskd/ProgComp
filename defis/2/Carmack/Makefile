WS_SRC		= src
WST_SRC		= tests
WS_TARGET	= ws
WST_TARGET	= wst

all: $(WS_TARGET) $(WST_TARGET)

$(WS_TARGET):
	@$(MAKE) -C $(WS_SRC)
	ln -s $(WS_SRC)/$(WS_TARGET) $(WS_TARGET)

$(WST_TARGET):
	@$(MAKE) -C tests
	ln -s $(WST_SRC)/$(WST_TARGET) $(WST_TARGET)

clean:
	@$(MAKE) clean -C $(WS_SRC)
	@$(MAKE) clean -C $(WST_SRC)

mrproper: clean
	@$(MAKE) mrproper -C $(WS_SRC)
	@$(MAKE) mrproper -C $(WST_SRC)
	rm -rf $(WS_TARGET) $(WST_TARGET)
