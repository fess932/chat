import neffos from 'neffos.js'

class Client {
  conn: neffos.Conn | undefined
  nsConn: neffos.NSConn | undefined
  constructor() {}

  async dial(host: string) {
    this.conn = await neffos.dial(host, {
      v1: {
        echo: echo,
      },
    })
  }

  async ping() {
    if (!this.conn) {
      console.log('no connection')
      return
    }
    this.nsConn = await this.conn.connect('v1')
    this.nsConn.emit('echo', 'Greetings!')
  }
}

function echo(nsConn: neffos.NSConn, msg: neffos.Message) {
  console.log('msg: ', msg)
}

export { Client }
