SHELL = /bin/sh
OCAMLBUILD = ocamlbuild
DEBUGFLAG = -g
OCAMLCFLAGS = -cflags $(DEBUGFLAG),-annot,-bin-annot,-w,A
OCAMLLFLAGS = -lflags $(DEBUGFLAG)
OCAMLLIBS =

RM = rm -f
TARGET = anagram.byte

build = $(OCAMLBUILD) $(OCAMLCFLAGS) $(OCAMLLFLAGS) $(OCAMLLIBS)
clean = $(OCAMLBUILD) -clean

.SUFFIXES:
.PHONY: all $(TARGET) clean

all: $(TARGET)
	ln -fs $(TARGET) anagram

$(TARGET):
	$(build) $@

clean:
	$(clean)
	$(RM) anagram
