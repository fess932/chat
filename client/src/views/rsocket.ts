import {
  RSocketClient,
  BufferEncoders,
  MESSAGE_RSOCKET_COMPOSITE_METADATA,
} from 'rsocket-core'

import RSocketWebSocketClient from 'rsocket-websocket-client'

const keepAlive = 60000
const lifetime = 180000
const dataMimeType = 'application/octet-stream'
const metadataMimeType = MESSAGE_RSOCKET_COMPOSITE_METADATA.string

const client = new RSocketClient({
  setup: {
    keepAlive,
    lifetime,
    dataMimeType,
    metadataMimeType,
  },
  transport: new RSocketWebSocketClient(
    {
      url: 'wss://localhost:7878',
      wsCreator: (url) => new WebSocket(url),
      debug: true,
    },
    BufferEncoders
  ),
})

export { client }
