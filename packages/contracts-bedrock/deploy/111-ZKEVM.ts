import { DeployFunction } from 'hardhat-deploy/dist/types'
import '@eth-optimism/hardhat-deploy-config'
import '@nomiclabs/hardhat-ethers'

import { deploy } from '../src/deploy-utils'

const deployFn: DeployFunction = async (hre) => {
  const l1 = hre.network.companionNetworks['l1']
  const deployConfig = hre.getDeployConfig(l1)

  await deploy({
    hre,
    name: 'ZKEVM',
    args: [],
  })
}

deployFn.tags = ['ZKEVM', 'l2']

export default deployFn
