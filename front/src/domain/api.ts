import { Configuration, ContestsApi, UsersApi, EntriesApi } from '../generated'

// TODO: ref env
export const apiBasePath = 'http://localhost:8081'

const baseConf = new Configuration({
  basePath: apiBasePath,
  credentials: 'include',
})

export const userClient = new UsersApi(baseConf)
export const contestClient = new ContestsApi(baseConf)
export const entryClient = new EntriesApi(baseConf)
