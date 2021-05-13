import { MessageBar, MessageBarType } from '@fluentui/react';
import React, { useEffect, useState } from 'react';

export interface IErrorMessageBar {
  error?: Error
}

export const ErrorMessageBar = (props: IErrorMessageBar) => {
  const { error } = props;

  const [_error, _setError] = useState<Error | undefined>();

  useEffect(() => {
    _setError(error);
  }, [error])

  return (<>
    {
      _error && <MessageBar styles={{ icon: { height: 16, lineHeight: '14px' } }} messageBarType={MessageBarType.blocked} onDismiss={() => { _setError(undefined); }} truncated={true}>
        {_error.message}
      </MessageBar>
    }
  </>)
}