SHELL = /bin/sh
RM = rm -f

OCAMLBUILD = ocamlbuild
OCAMLCFLAGS = -cflags -annot,-bin-annot,-safe-string,-w,A

TARGET = anagram

build = $(OCAMLBUILD) $(OCAMLCFLAGS)
clean = $(OCAMLBUILD) -clean

.PHONY: all $(TARGET) clean

all: $(TARGET)
	@ln -fs $<.native $<

$(TARGET):
	$(build) $@.native

clean:
	$(clean)
	@$(RM) anagram
