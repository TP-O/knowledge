var m = undefined;

/**
 * Definition for a binary tree node.
 * function TreeNode(val, left, right) {
 *     this.val = (val===undefined ? 0 : val)
 *     this.left = (left===undefined ? null : left)
 *     this.right = (right===undefined ? null : right)
 * }
 */
/**
 * @param {TreeNode} root
 * @return {number}
 */
var maxPathSum = function(root) {
    m = undefined;
    calculatePath(root, 0);
    return m;
};

var calculatePath = function (node, path) {
    node.path = path + node.val;

    if (!(m > node.val)) {
        m = node.val;
    }

    if (!(m > node.path)) {
        m = node.path;
    }

    if (!node.left && !node.right) {
        if (!(m > node.path)) {
            m = node.path;
        }

        return;
    }    

    var completedLeftPath = undefined;
    var completedRightPath = undefined;

    if (node.left) {
        var leftPath = undefined;
        var rightPath = undefined;

        var parentPath = node.path + node.left.val;
        if (!(m > parentPath)) {
            m = parentPath;
        }

        if (!(m > node.left.val)) {
            m = node.left.val;
        }

        if (node.path < node.val && !(m > node.val + node.left.val)) {
            m = node.val + node.left.val
        }

        if (node.left.left) {
            if (node.path < node.val) {
                calculatePath(node.left.left, node.val);
                node.left.left.path += path + node.left.val;
            } else if (parentPath < node.left.val) {
                calculatePath(node.left.left, node.left.val);
                node.left.left.path += node.path;
            } else { 
                calculatePath(node.left.left, parentPath);
            }

            if (node.left.left.path < node.left.path) {
                leftPath = node.left.path;
            } else {
                leftPath = node.left.left.path;
            }

            if (!(m > leftPath)) {
                m = leftPath;
            }
        }

        if (node.left.right) {
            if (node.path < node.val) {
                calculatePath(node.left.right, node.val);
                node.left.right.path += path + node.left.val;
            } else if (parentPath < node.left.val) {
                calculatePath(node.left.right, node.left.val);
                node.left.right.path += node.path;
            } else { 
                calculatePath(node.left.right, parentPath);
            }

            if (node.left.right.path < node.left.path) {
                rightPath = node.left.path;
            } else {
                rightPath = node.left.right.path;
            }

            if (!(m > rightPath)) {
                m = rightPath;
            }
        }

        // console.log(node.val, node.left?.val, node.right?.val, leftPath, rightPath);

        if (leftPath === undefined && rightPath !== undefined && rightPath > parentPath) {
            completedLeftPath = rightPath;

            if (!(m > rightPath - parentPath + node.left.val)) {
                m = rightPath - parentPath + node.left.val;
            }
        } else if (leftPath !== undefined && rightPath === undefined && leftPath > parentPath) {
            completedLeftPath = leftPath;

            if (!(m > leftPath - parentPath + node.left.val)) {
                m = leftPath - parentPath + node.left.val;
            }
        } else if (leftPath !== undefined && rightPath !== undefined && (rightPath > parentPath || leftPath > parentPath)) {
            if (leftPath > rightPath) {
                completedLeftPath = leftPath;
            } else {
                completedLeftPath = rightPath;
            }

            if (!(m > leftPath + rightPath - 2*parentPath + node.left.val)) {
                m = leftPath + rightPath - 2*parentPath + node.left.val;
            }
        } else if (parentPath < node.path) {
            completedLeftPath = node.path;
        } else {
            completedLeftPath = parentPath;
        }
    }

    // console.log(node.val, node.path, completedLeftPath)

    if (node.right) {
        var leftPath = undefined;
        var rightPath = undefined;

        var parentPath = node.path + node.right.val;
        if (!(m > parentPath)) {
            m = parentPath;
        }

        if (!(m > node.right.val)) {
            m = node.right.val;
        }

        if (node.path < node.val && !(m > node.val + node.right.val)) {
            m = node.val + node.right.val
        }

        if (node.right.left) {
            if (node.path < node.val) {
                calculatePath(node.right.left, node.val);
                node.right.left.path += path + node.right.val;
            } else if (parentPath < node.right.val) {
                calculatePath(node.right.left, node.right.val);
                node.right.left.path += node.path;
            } else { 
                calculatePath(node.right.left, parentPath);
            }

            if (node.right.left.path < node.right.path) {
                leftPath = node.right.path;
            } else {
                leftPath = node.right.left.path;
            }

            if (!(m > leftPath)) {
                m = leftPath;
            }
        }

        if (node.right.right) {
            if (node.path < node.val) {
                calculatePath(node.right.right, node.val);
                node.right.right.path += path + node.right.val;
            } else if (parentPath < node.right.val) {
                calculatePath(node.right.right, node.right.val);
                node.right.right.path += node.path;
            } else { 
                calculatePath(node.right.right, parentPath);
            }

            if (node.right.right.path < node.right.path) {
                rightPath = node.right.path;
            } else {
                rightPath = node.right.right.path;
            }

            if (!(m > rightPath)) {
                m =  rightPath;
            }
        }

        // console.log(leftPath, rightPath);

        if (leftPath === undefined && rightPath !== undefined && rightPath > parentPath) {
            completedRightPath = rightPath;

            if (!(m > rightPath - parentPath + node.right.val)) {
                m = rightPath - parentPath + node.right.val;
            }
        } else if (leftPath !== undefined && rightPath === undefined && leftPath > parentPath) {
            completedRightPath = leftPath;

            if (!(m > leftPath - parentPath + node.right.val)) {
                m = leftPath - parentPath + node.right.val;
            }
        } else if (leftPath !== undefined && rightPath !== undefined && (rightPath > parentPath || leftPath > parentPath)) {
            if (leftPath > rightPath) {
                completedRightPath = leftPath;
            } else {
                completedRightPath = rightPath;
            }

            if (!(m > leftPath + rightPath - 2*parentPath + node.right.val)) {
                m = leftPath + rightPath - 2*parentPath + node.right.val;
            }
        } else if (parentPath < node.path) {
            completedRightPath = node.path;
        } else {
            completedRightPath = parentPath;
        }
    }

    // console.log(node.val, completedRightPath)

    // console.log(completedLeftPath, completedRightPath)

    if (completedLeftPath === undefined && completedRightPath !== undefined && completedRightPath > node.path) {
        node.path = completedRightPath;
    } else if (completedLeftPath !== undefined && completedRightPath === undefined && completedLeftPath > node.path) {
        node.path = completedLeftPath;
    } else if (completedLeftPath !== undefined && completedRightPath !== undefined && (completedLeftPath > node.path || completedRightPath > node.path)) {
        console.log(node.val, completedLeftPath, completedRightPath, node.path, node.val);

        if (!(m > completedLeftPath + completedRightPath - 2*node.path + node.val)) {
            m = completedLeftPath + completedRightPath - 2*node.path + node.val;
        }

        if (completedLeftPath > completedRightPath) {
            node.path = completedLeftPath;
        } else {
            node.path = completedRightPath;
        }
    } else if (node.path < path) {
        node.path = path;
    }
}
