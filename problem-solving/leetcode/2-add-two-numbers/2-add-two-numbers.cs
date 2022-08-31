/**
 * Definition for singly-linked list.
 * public class ListNode {
 *     public int val;
 *     public ListNode next;
 *     public ListNode(int val=0, ListNode next=null) {
 *         this.val = val;
 *         this.next = next;
 *     }
 * }
 */
public class Solution {    
    public ListNode AddTwoNumbers(ListNode l1, ListNode l2) {
        ListNode currentL1Node = l1;
        ListNode currentL2Node = l2;
        int sum = currentL1Node.val + currentL2Node.val;
        int tensDigit = sum / 10;
        int unitDigit = sum % 10;
        ListNode resultList = new ListNode(unitDigit);
        ListNode currentResultNode = resultList;
        
        while (currentL1Node.next != null && currentL2Node.next != null) {
            currentL1Node = currentL1Node.next;
            currentL2Node = currentL2Node.next;
            
            sum = currentL1Node.val + currentL2Node.val + tensDigit;
            tensDigit = sum / 10;
            unitDigit = sum % 10;
            
            ListNode newNode = new ListNode(unitDigit);
            currentResultNode.next = newNode;
            currentResultNode = currentResultNode.next;
        }
        
        while (currentL1Node.next != null) {
            currentL1Node = currentL1Node.next;
            sum = currentL1Node.val + tensDigit;
            tensDigit = sum / 10;
            unitDigit = sum % 10;
            
            ListNode newNode = new ListNode(unitDigit);
            currentResultNode.next = newNode;
            currentResultNode = currentResultNode.next;
        }
        
        while (currentL2Node.next != null) {
            currentL2Node = currentL2Node.next;
            sum = currentL2Node.val + tensDigit;
            tensDigit = sum / 10;
            unitDigit = sum % 10;
            
            ListNode newNode = new ListNode(unitDigit);
            currentResultNode.next = newNode;
            currentResultNode = currentResultNode.next;
        }
        
        if (tensDigit != 0) {
            ListNode newNode = new ListNode(tensDigit);
            currentResultNode.next = newNode;
            currentResultNode = currentResultNode.next;
        }
        
        return resultList;
    }
}
