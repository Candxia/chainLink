import service from '../index'

export const blockchainApi = {
  getBlockList(params) {
    return service({ url: '/blockchain/block/list', method: 'get', params })
  },
  getTransactionList(params) {
    return service({ url: '/blockchain/transaction/list', method: 'get', params })
  },
  getContractList(params) {
    return service({ url: '/blockchain/contract/list', method: 'get', params })
  },
  deployContract(data) {
    return service({ url: '/blockchain/contract/deploy', method: 'post', data })
  },
}

