{
  description = "Interact with Zube from the CL";

  # Nixpkgs / NixOS version to use.
  inputs.nixpkgs.url = "nixpkgs/nixos-unstable";

  inputs.flake-compat = {
    url = "github:edolstra/flake-compat";
    flake = false;
  };

  outputs = { self, nixpkgs, ... }:
    let
      version = "0.3.3";

      # System types to support.
      supportedSystems =
        [ "x86_64-linux" "x86_64-darwin" "aarch64-linux" "aarch64-darwin" ];

      # Helper function to generate an attrset '{ x86_64-linux = f "x86_64-linux"; ... }'.
      forAllSystems = nixpkgs.lib.genAttrs supportedSystems;

      # Nixpkgs instantiated for supported system types.
      nixpkgsFor = forAllSystems (system: import nixpkgs { inherit system; });

    in {

      # Provide some binary packages for selected system types.
      packages = forAllSystems (system:
        let pkgs = nixpkgsFor.${system};
        in rec {
          zube-cli = pkgs.buildGoModule {
            pname = "zube-cli";
            inherit version;

            src = ./.;

            vendorHash =
              "sha256-SHZxuwpyOr97NspJ8wbKPziuaTKAtUmNHGFTkkvZQFc=";

            checkPhase = ''
              make test
            '';
          };
          default = zube-cli;
        });

      apps = forAllSystems (system: rec {
        zube-cli = {
          type = "app";
          program = "${self.packages.${system}.go-hello}/bin/zube-cli";
        };
        default = zube-cli;
      });

      # The default package for 'nix build'. This makes sense if the
      # flake provides only one package or there is a clear "main"
      # package.
      defaultPackage = forAllSystems (system: self.packages.${system}.default);

      defaultApp = forAllSystems (system: self.apps.${system}.default);

      devShells = forAllSystems (system:
        let pkgs = nixpkgsFor.${system};
        in {
          default = pkgs.mkShell {
            buildInputs = with pkgs; [ go gopls goimports go-tools ];
          };
        });

      devShell = forAllSystems (system: self.devShells.${system}.default);
    };
}
