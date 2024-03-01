with import <nixpkgs> { };

buildGoModule rec {
  pname = "zube-cli";
  version = "0.3.3";
  binary = "zube";

  src = fetchFromGitHub {
    owner = "platogo";
    repo = "zube-cli";
    rev = "${version}";
    hash = "sha256-tF4cxqZ/9I+USrUxmGeRc7VxnPVr4vrObmXthYGNDoA=";
  };

  vendorHash = "sha256-CZHEYjy1oL4ns0MEDE3JxDfeUoTaX+JZktB0EpqVIPk=";

  ldflags =
    [ "-s" "-w" "-X main.Version=${version}" "-X main.Commit=${version}" ];

  meta = with lib; {
    description = "Interact with Zube from the CLI";
    homepage = "https://github.com/platogo/zube-cli";
    license = licenses.gpl3;
    maintainers = with maintainers; [ danirukun ];
  };

  installPhase = ''
    runHook preInstall
    install -D $GOPATH/bin/${pname} $out/bin/zube # Rename binary here
    runHook postInstall
  '';

  checkPhase = ''
    make test
  '';
}
