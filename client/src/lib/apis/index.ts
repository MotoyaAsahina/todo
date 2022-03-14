import { Apis, Configuration } from './generated'

export const apis = new Apis(
  new Configuration({
    basePath: '/api'
  })
)

export * from './generated'
