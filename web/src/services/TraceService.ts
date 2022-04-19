import {AssertionResult, AssertionResultList, ITrace, spanAssertionResult, Test, TestAssertionResult} from '../types';
import {runAssertionByTrace} from './AssertionService';

export const runTest = (trace: ITrace, {assertions = []}: Test) => {
  const resultList = assertions?.map(assertion => runAssertionByTrace(trace, assertion));

  return resultList;
};

export const parseTestResultToAssertionResultList = (
  assertionResult: AssertionResultList,
  {assertions}: Test,
  {resourceSpans}: ITrace
): AssertionResult[] => {
  return assertionResult.map(({assertionId, spanAssertionResults}) => {
    const assertion = assertions.find(({assertionId: id}) => id === assertionId);

    return {
      assertion: assertion!,
      spanListAssertionResult: spanAssertionResults.map(({spanId, passed, observedValue, spanAssertionId}) => {
        const resourceSpan = resourceSpans.find(({instrumentationLibrarySpans}) => {
          const span = instrumentationLibrarySpans.find(({spans}) => {
            return spans.find(({spanId: id}) => id === spanId);
          });

          return span;
        });

        const spanAssertion = assertion?.spanAssertions.find(({spanAssertionId: id}) => id === spanAssertionId);

        return {
          span: resourceSpan!,
          resultList: spanAssertion ? [{...spanAssertion, spanId, hasPassed: passed, actualValue: observedValue}] : [],
        };
      }),
    };
  });
};

export const parseAssertionResultListToTestResult = (assertionResultList: AssertionResult[]): TestAssertionResult => {
  const {totalFailedCount} = getTestResultCount(assertionResultList);

  return {
    assertionResultState: !totalFailedCount,
    assertionResult: assertionResultList.map(({assertion, spanListAssertionResult}) => ({
      assertionId: assertion.assertionId,
      spanAssertionResults: spanListAssertionResult.reduce<spanAssertionResult[]>(
        (accList, {resultList}) =>
          accList.concat(
            resultList.map(({spanId, hasPassed, actualValue, spanAssertionId = ''}) => ({
              spanAssertionId,
              spanId,
              passed: hasPassed,
              observedValue: actualValue,
            }))
          ),
        []
      ),
    })),
  };
};

export const getTestResultCount = (assertionResultList: AssertionResult[]) => {
  const [totalPassedCount, totalFailedCount] = assertionResultList.reduce<[number, number]>(
    ([innerTotalPassedCount, innerTotalFailedCount], {spanListAssertionResult}) => {
      const [passed, failed] = spanListAssertionResult.reduce<[number, number]>(
        ([passedResultCount, failedResultCount], {resultList}) => {
          const passedCount = resultList.filter(({hasPassed}) => hasPassed).length;
          const failedCount = resultList.filter(({hasPassed}) => !hasPassed).length;

          return [passedResultCount + passedCount, failedResultCount + failedCount];
        },
        [0, 0]
      );

      return [innerTotalPassedCount + passed, innerTotalFailedCount + failed];
    },
    [0, 0]
  );

  return {totalFailedCount, totalPassedCount};
};
