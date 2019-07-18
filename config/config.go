package config

type Server struct {
	host string `yaml:"host"`
	port int    `yaml:"port"`
}

type Blog struct {
	articleDir string `yaml:"article_dir"`
	flushTime  int    `yaml:"flush_time"`
	title      string `yaml:"title"`
}

type Conf struct {
	server Server `yaml:"Server"`
	blog   Blog   `yaml:"Blog"`
}
