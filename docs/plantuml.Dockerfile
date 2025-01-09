FROM openjdk:17-jdk-slim

# curl のインストール
RUN apt-get update && apt-get install -y curl

# 日本語フォントとgraphviz のインストール
RUN apt-get update && apt-get install -y fonts-ipafont-gothic fonts-ipafont-mincho fonts-noto-cjk graphviz && rm -rf /var/lib/apt/lists/*

# PlantUML のダウンロード
WORKDIR /app
RUN curl -L https://github.com/plantuml/plantuml/releases/download/v1.2024.8/plantuml-1.2024.8.jar -o plantuml.jar

# 作業ディレクトリの設定とpumlディレクトリ作成
WORKDIR /work
RUN mkdir puml

# puml ファイルをコピー
COPY puml ./puml

# SVG 生成スクリプト (entrypoint.sh)
COPY puml.sh .
RUN chmod +x puml.sh

# エントリーポイントを設定
ENTRYPOINT ["./puml.sh"]