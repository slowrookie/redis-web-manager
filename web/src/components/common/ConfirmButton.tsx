import { DefaultButton, Dialog, DialogFooter, IconButton, IDialogContentProps, PrimaryButton, TooltipHost } from '@fluentui/react';
import React, { useEffect, useState } from 'react';
import { useTranslation } from 'react-i18next';

export interface IConfirmButtonProps {
  label: string
  dialogContentProps: IDialogContentProps
  disabled: boolean
  onConfirm: () => void
  onCancel?: () => void
}

export const ConfirmButton = (props: IConfirmButtonProps) => {
  const { label, dialogContentProps, onConfirm, onCancel } = props;

  const { t } = useTranslation();

  const [hiddenDialog, setHiddenDialog] = useState(true),
    [disabled, setDisabled] = useState(props.disabled);

  useEffect(() => {
    setDisabled(props.disabled);
  }, [props.disabled])

  return (<div>
    <TooltipHost content={label}>
      <IconButton iconProps={{ iconName: 'blocked2', style: { height: 'auto' } }}
        disabled={disabled}
        onClick={() => {
          setHiddenDialog(false);
        }} />
    </TooltipHost>
    <Dialog styles={{ main: { minHeight: 'auto' } }} hidden={hiddenDialog} dialogContentProps={dialogContentProps}
      onDismiss={() => setHiddenDialog(true)}>
      <DialogFooter>
        <PrimaryButton onClick={() => { setHiddenDialog(true); onConfirm(); }} text={t('OK')} />
        <DefaultButton onClick={() => { setHiddenDialog(true); onCancel && onCancel(); }} text={t('Cancel')} />
      </DialogFooter>
    </Dialog>
  </div>)
}