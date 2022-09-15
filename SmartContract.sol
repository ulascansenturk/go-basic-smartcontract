// SPDX-License-Identifier: GPL-3.0
pragma solidity >=0.7.0 <0.9.0;

contract MySmartContract {
    function sortArray(uint[] memory arr) public pure returns (uint[] memory) {
        uint[] memory sortedArray = arr;
        for (uint i = 0; i < sortedArray.length; i++) {
            for (uint j = i + 1; j < sortedArray.length; j++) {
                if (sortedArray[i] > sortedArray[j]) {
                    uint temp = sortedArray[i];
                    sortedArray[i] = sortedArray[j];
                    sortedArray[j] = temp;
                }
            }
        }
        return sortedArray;
    }



}
