import { MessageBar, MessageBarType } from '@fluentui/react';
import React, { useEffect, useState } from 'react';

export interface IErrorMessageBar {
  error: string | undefined
}

export const ErrorMessageBar = (props: IErrorMessageBar) => {
  const { error } = props;

  const [_error, _setError] = useState<string>();

  useEffect(() => {
    _setError(error);
  }, [error])

  return (<>
    {
      _error && <MessageBar styles={{ icon: { height: 16, lineHeight: '14px' } }} messageBarType={MessageBarType.blocked} onDismiss={() => { _setError(""); }} truncated={true}>
        {_error}
      </MessageBar>
    }
  </>)
}