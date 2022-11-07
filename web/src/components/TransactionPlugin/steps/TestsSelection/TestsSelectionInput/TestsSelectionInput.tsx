import {closestCenter, DndContext, KeyboardSensor, PointerSensor, useSensor, useSensors} from '@dnd-kit/core';
import {arrayMove, SortableContext, sortableKeyboardCoordinates, verticalListSortingStrategy} from '@dnd-kit/sortable';
import {Col, Row, Select} from 'antd';
import {noop} from 'lodash';
import {useCallback, useEffect, useState} from 'react';
import {TTest} from 'types/Test.types';
import TestItemList from './TestItemList';

interface IProps {
  value?: string[];
  testList: TTest[];
  onChange?(tests: string[]): void;
}

export interface ISortableTest {
  id: string;
  test: TTest;
}

const TestsSelectionInput = ({value = [], onChange = noop, testList}: IProps) => {
  const [selectedTestList, setSelectedTestList] = useState<ISortableTest[]>([]);
  const onSelectedTest = useCallback(
    (testId: string) => {
      onChange([...value, testId]);
    },
    [onChange, value]
  );

  useEffect(() => {
    if (testList.length) {
      setSelectedTestList(
        value.map((testId, index) => {
          const test = testList.find(({id}) => id === testId)!;

          return {
            test,
            id: `${test.id}-${index}`,
          };
        })
      );
    }
  }, [testList, value]);

  const sensors = useSensors(
    useSensor(PointerSensor),
    useSensor(KeyboardSensor, {
      coordinateGetter: sortableKeyboardCoordinates,
    })
  );

  const onSortEnd = useCallback(
    ({active, over}) => {
      if (active.id !== over.id) {
        const oldIndex = selectedTestList.findIndex(({id}) => id === active.id);
        const newIndex = selectedTestList.findIndex(({id}) => id === over.id);
        const updatedList = arrayMove(selectedTestList, oldIndex, newIndex);

        setSelectedTestList(updatedList);
        onChange(updatedList.map(({test}) => test.id));
      }
    },
    [onChange, selectedTestList]
  );

  const onDelete = useCallback(
    (testId: string) => {
      onChange(value.filter(id => id !== testId));
      setSelectedTestList(selectedTestList.filter(test => test.id !== testId));
    },
    [onChange, selectedTestList, value]
  );

  return (
    <>
      <DndContext sensors={sensors} collisionDetection={closestCenter} onDragEnd={onSortEnd}>
        <SortableContext items={selectedTestList} strategy={verticalListSortingStrategy}>
          <TestItemList items={selectedTestList} onDelete={onDelete} />
        </SortableContext>
      </DndContext>
      <Row gutter={12}>
        <Col span={18}>
          <span>
            <Select
              placeholder="Add a test"
              onChange={onSelectedTest}
              value={null}
              showSearch
              filterOption={(input, option) =>
                (option!.children as unknown as string).toLowerCase().includes(input.toLowerCase())
              }
            >
              {testList.map(({id, name}) => (
                <Select.Option value={id} key={id}>
                  {name}
                </Select.Option>
              ))}
            </Select>
          </span>
        </Col>
      </Row>
    </>
  );
};

export default TestsSelectionInput;
