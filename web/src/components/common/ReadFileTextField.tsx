import { CommandBarButton, ITextFieldProps, Stack, TextField } from '@fluentui/react';
import React, { useRef, useState } from 'react';
import { useTranslation } from 'react-i18next';
import {readFile} from '../../services/file.service';

export const ReadFileTextFiled = (props: ITextFieldProps) => {
  const { t } = useTranslation();
  const fileChosenRef = useRef<any>();
  const [errorMessage, setErrorMessage] = useState<string | undefined>();

  return <TextField {...props} errorMessage={errorMessage} onRenderLabel={(fieldProps, defaultRender: any) => {
    return (
      <Stack horizontal>
        {defaultRender(fieldProps)}
        <Stack.Item grow={1}><span></span></Stack.Item>
        <input type="file" style={{ display: 'none' }} ref={fileChosenRef} onChange={(e) => {
          if (!e.target.files) return;
          let file = e.target.files[0];
          const fileReader = new FileReader()
          fileReader.onloadend = () => {
            props.onChange && props.onChange(e, fileReader.result?.toString());
          }
          fileReader.readAsText(file);
        }} />
        <CommandBarButton iconProps={{ iconName: 'Upload' }} text={t('Upload')} onClick={(e: any) => {
        //   fileChosenRef.current.click();
          readFile().then(v => {
              props.onChange && props.onChange(e, v);
          }).catch(err => setErrorMessage(err))
        }}></CommandBarButton>
      </Stack>
    )
  }} />
}
