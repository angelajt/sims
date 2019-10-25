# go through all the sims

TOPTARGETS := all clean mac linux windows

SUBDIRS := $(wildcard */.)

$(TOPTARGETS): $(SUBDIRS)
$(SUBDIRS):
	$(MAKE) -C $@ $(MAKECMDGOALS)

.PHONY: $(TOPTARGETS) $(SUBDIRS)

# NOTE: MUST update version number here prior to running 'make release' and edit this file! 
VERS=v0.5.5
PACKAGE=sims
GIT_COMMIT=`git rev-parse --short HEAD`
VERS_DATE=`date -u +%Y-%m-%d\ %H:%M`
VERS_FILE=version.go

release:
	# /bin/rm -f $(VERS_FILE)
	# @echo "// WARNING: auto-generated by Makefile release target -- run 'make release' to update" > $(VERS_FILE)
	# @echo "" >> $(VERS_FILE)
	# @echo "package $(PACKAGE)" >> $(VERS_FILE)
	# @echo "" >> $(VERS_FILE)
	# @echo "const (" >> $(VERS_FILE)
	# @echo "	Version = \"$(VERS)\"" >> $(VERS_FILE)
	# @echo "	GitCommit = \"$(GIT_COMMIT)\" // the commit JUST BEFORE the release" >> $(VERS_FILE)
	# @echo "	VersionDate = \"$(VERS_DATE)\" // UTC" >> $(VERS_FILE)
	# @echo ")" >> $(VERS_FILE)
	# @echo "" >> $(VERS_FILE)
	# /bin/cat $(VERS_FILE)
	git commit -am "$(VERS) release"
	git tag -a $(VERS) -m "$(VERS) release"
	git push
	git push origin --tags

