// SPDX-License-Identifier: UNLICENSED

pragma solidity ^0.8.0;
import "@openzeppelin/contracts/token/ERC1155/ERC1155.sol";

contract Goods is ERC1155 {
    uint256 public successNum;

    constructor() ERC1155("") {}

    function batchMint(address owner, uint256[] memory ids) public {
        uint256[] memory amounts = new uint256[](ids.length);

        for (uint256 i = 0; i < ids.length; ++i) {
            amounts[i] = 1;
            successNum++;
        }
        _mintBatch(owner, ids, amounts, "");
    }

    function mint(address to, uint256 id) public {
        successNum++;
        _mint(to, id, 1, "");
    }

    function batchTransfer(
        address from,
        address to,
        uint256[] memory ids
    ) public {
        uint256[] memory amounts = new uint256[](ids.length);

        for (uint256 i = 0; i < ids.length; ++i) {
            amounts[i] = 1;
        }
        safeBatchTransferFrom(from, to, ids, amounts, "");
    }

    function transfer(
        address from,
        address to,
        uint256 id
    ) public {
        safeTransferFrom(from, to, id, 1, "");
    }

    function getSuccessNum() public view returns (uint256) {
        return successNum;
    }
}
