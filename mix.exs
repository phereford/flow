defmodule Flow.MixProject do
  use Mix.Project

  def project do
    [
      app: :flow,
      version: "0.1.6",
      elixir: "~> 1.11",
      start_permanent: Mix.env() == :prod,
      deps: deps()
    ]
  end

  # Run "mix help compile.app" to learn about applications.
  def application do
    [
      extra_applications: [:logger]
    ]
  end

  # Run "mix help deps" to learn about dependencies.
  defp deps do
    [
      {:grpc, github: "elixir-grpc/grpc"},
      {:protobuf, "~> 0.7.1"},
      {:google_protos, "~> 0.1"}
    ]
  end
end
