Name:           grasshopper
Version:        0.0.24
Release:        1%{?dist}
Summary:        This will make a Nulecule GO!

License:        LGPLv3+
URL:            https://github.com/goern/grasshopper
Source0:        grasshopper-%{version}.tar.gz

ExclusiveArch:  x86_64
BuildRequires:	golang-bin git rubygem-asciidoctor
Requires:       golang

%ifarch x86_64
  %global GOARCH amd64
%endif

%description
This will make a Nulecule GO!

%prep
%setup -q -n %{name}-%{version}

mkdir -p $RPM_BUILD_ROOT/%{_bindir}

%build
GOPATH="$(pwd)"
GOBIN="$GOPATH/bin"
GOOS=linux
GOARCH="%{GOARCH}"
export GOPATH GOBIN GOOS GOARCH

LC_ALL=C PATH="$PATH:$GOBIN" go get github.com/tools/godep
LC_ALL=C PATH="$PATH:$GOBIN" go get github.com/goern/grasshopper
LC_ALL=C PATH="$PATH:$GOBIN" GRASSHOPPER_VERSION=%{version} make
LC_ALL=C PATH="$PATH:$GOBIN" GRASSHOPPER_VERSION=%{version} make doc
cp grasshopper-%{version} $RPM_BUILD_ROOT/%{_bindir}/grasshopper-%{version}

%clean
[ "$RPM_BUILD_ROOT" != "/" ] && rm -rf $RPM_BUILD_ROOT

%files
%defattr(0644,root,root,0755)

%attr(0755,-,-) %{_bindir}/grasshopper-%{version}

%doc AUTHORS
%doc LICENSE
%doc README.html

%post
alternatives --install %{_bindir}/grasshopper grasshopper %{_bindir}/grasshopper-{version} %{alternatives_priority}

%preun
alternatives --remove grasshopper %{_bindir}/grasshopper-%{version}

%changelog
* Sun Nov 08 2015 Christoph Görn <goern@redhat.com> 0.0.24-1
- 

* Sun Nov 08 2015 Christoph Görn <goern@redhat.com> 0.0.23-1
-

* Sun Nov 08 2015 Christoph Görn <goern@redhat.com> 0.0.22-1
- update to build with GO 1.5.1 (goern@redhat.com)

* Sun Nov 08 2015 Christoph Görn <goern@redhat.com> 0.0.21-1
- fix the copr releaser (goern@redhat.com)
- change from Markdown to Asciidoc (goern@redhat.com)

* Sun Nov 08 2015 Christoph Görn <goern@redhat.com>
- fix the copr releaser (goern@redhat.com)
- change from Markdown to Asciidoc (goern@redhat.com)

* Sun Nov 08 2015 Christoph Görn <goern@redhat.com> 0.0.20-1
- fix some more build path issues (goern@redhat.com)

* Wed Nov 04 2015 Christoph Görn <goern@redhat.com> 0.0.12-1
- add git as a build requirement, due to `go get` (goern@redhat.com)

* Wed Nov 04 2015 Christoph Görn <goern@redhat.com> 0.0.9-1
- add godep as a build step (goern@redhat.com)

* Wed Nov 04 2015 Christoph Görn <goern@redhat.com> 0.0.8-1
- fix a type in .spec (goern@redhat.com)

* Wed Nov 04 2015 Christoph Görn <goern@redhat.com> 0.0.7-1
- new package built with tito

* Wed Nov 04 2015 Christoph Görn <goern@redhat.com>
- move to rel-eng/ (goern@redhat.com)
- add man page (stub) (goern@redhat.com)
- remove the .srpm (goern@redhat.com)

* Tue Nov 03 2015 Christoph Görn <goern@redhat.com> 0.0.5-1
- new package built with tito

* Tue Nov 03 2015 Christoph Görn
- initial RPMification
