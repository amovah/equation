package equation

type takenDetail struct {
	taken      bool
	takenBy    uint
	level      uint
	precedence uint16
}

func cleanSeenNode(seen *map[uint]takenDetail, graph *equationGraph, id uint) error {
	node, err := graph.getNode(id)
	if err != nil {
		return err
	}

	if node.equationNodeType() == graphNumberNode {
		(*seen)[id] = takenDetail{}
		return nil
	}

	operatorNode, err := graph.getOperatorNode(id)
	if err != nil {
		return err
	}

	(*seen)[id] = takenDetail{}

	for _, operatorArg := range operatorNode.operatorArgs {
		(*seen)[operatorArg] = takenDetail{}
	}

	return nil
}

func canTakeThisOperator(
	seen *map[uint]takenDetail,
	graph *equationGraph,
	level uint,
	precedence uint16,
	subjectNodeId uint,
) (bool, error) {
	if (*seen)[subjectNodeId].level > level || ((*seen)[subjectNodeId].level == level && (*seen)[subjectNodeId].precedence > precedence) {
		return false, nil
	}

	node, err := graph.getNode(subjectNodeId)
	if err != nil {
		return false, err
	}

	if node.equationNodeType() == graphOperatorNode {
		operatorNode, err := graph.getOperatorNode(subjectNodeId)
		if err != nil {
			return false, err
		}

		for _, operatorArg := range operatorNode.operatorArgs {
			canTake, err := canTakeThisOperator(seen, graph, level, precedence, operatorArg)
			if err != nil {
				return false, err
			}

			if canTake == false {
				return false, nil
			}
		}
	}

	return true, nil
}

func solve(graph equationGraph) error {
	seen := make(map[uint]takenDetail)

	for v := range graph.chanOverTree() {
		if v.equationNodeType() == graphOperatorNode {
			node, err := graph.getOperatorNode(v.equationId())
			if err != nil {
				return err
			}

			skip := false
			for _, arg := range node.operatorArgs {
				canTake, err := canTakeThisOperator(
					&seen,
					&graph,
					node.graphLevel,
					node.operator.precedence,
					arg,
				)
				if err != nil {
					return err
				}

				if !canTake {
					skip = true
					break
				}
			}

			if skip {
				continue
			}

			if seen[node.id].taken {
				err := cleanSeenNode(&seen, &graph, seen[node.id].takenBy)
				if err != nil {
					return nil
				}
			}

			for _, arg := range node.operatorArgs {
				if seen[arg].taken {
					err := cleanSeenNode(&seen, &graph, seen[arg].takenBy)
					if err != nil {
						return err
					}
				}

				seen[arg] = takenDetail{
					taken:      true,
					takenBy:    node.id,
					level:      node.graphLevel,
					precedence: node.operator.precedence,
				}
			}

			seen[node.id] = takenDetail{
				taken:      true,
				takenBy:    node.id,
				level:      node.graphLevel,
				precedence: node.operator.precedence,
			}
		}
	}

	return nil
}
