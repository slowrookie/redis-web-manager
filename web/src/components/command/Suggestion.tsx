import { Callout, CheckboxVisibility, DetailsList, DetailsListLayoutMode, DirectionalHint, FocusZone, FocusZoneDirection, SelectionMode, TextField } from '@fluentui/react';
import { useId } from '@fluentui/react-hooks';
import React, { ChangeEvent, KeyboardEvent, useEffect, useState } from 'react';
import { Connection, suggestions } from '../../services/connection.service';

export interface ISuggestionProps {
  connection: Connection,
  onChange?: (command?: string) => void
  onSearch: (command?: string) => void
  errorMessage?: string
}


export const Suggestion = (props: ISuggestionProps) => {

  const [command, setCommand] = useState<string>(),
    searchBoxId = useId(`command-callout-button-${props.connection.id}`),
    [isCalloutVisible, setCalloutVisible] = useState(false),
    [expects, setExpects] = useState<Array<string | number>>([]),
    [_errorMessage, _setErrorMessage] = useState<string>();

  const handleChange = (e?: ChangeEvent<any>, value?: string) => {
    setCommand(value || "");
    _setErrorMessage("");
    
    if (value) {
      suggestions({ id: props.connection.id, commands: [[value]] }).then((v: any) => {
        if (v) {
          const tokens = value.split(' ');
          const matchToken = tokens[tokens.length - 1];
          const eps = matchToken ? (v as Array<string>).filter(e => e.indexOf(matchToken.toUpperCase()) >= 0) : v;
          setExpects(eps);
          setCalloutVisible(!!eps.length);
        } else {
          setExpects([])
          setCalloutVisible(false);
        }
      })
    } else {
      setCalloutVisible(false);
    }
  }

  const handleItemInvoked = (item?: any, index?: number, ev?: Event) => {
    if (command) {
      const cmdSplits = command.split(' ');
      cmdSplits[cmdSplits.length - 1] = item;
      setCommand(cmdSplits.join(' ').trim());
    } else {
      setCommand(item.trim());
    }
    setCalloutVisible(false);
  }

  const handleKeyUp = (e: KeyboardEvent) => {
    if (e.code === 'Enter') {
      props.onSearch && props.onSearch(command);
    }
  }

  useEffect(() => {
    _setErrorMessage(props.errorMessage);
  }, [props.errorMessage])

  useEffect(() => {
    props.onChange && props.onChange(command);
  }, [command, props])

  return <>
    <div style={{ width: '100%' }} >
      <TextField
        id={searchBoxId}
        styles={{ root: { height: 24, paddingLeft: 0 }, fieldGroup: { height: 24 }, field: { padding: 0 } }}
        autoFocus
        underlined={true}
        value={command}
        onChange={handleChange}
        onKeyUp={handleKeyUp}
        errorMessage={_errorMessage}
      />
      <Callout id='SuggestionContainer'
        ariaLabelledBy={'callout-suggestions'}
        gapSpace={5}
        coverTarget={false}
        alignTargetEdge={true}
        onDismiss={ev => setCalloutVisible(false)}
        setInitialFocus={false}
        hidden={!isCalloutVisible}
        calloutMaxHeight={300}
        style={{ width: 295 }}
        target={`#${searchBoxId}`}
        directionalHint={DirectionalHint.bottomLeftEdge}
        isBeakVisible={false}
      >
        <FocusZone direction={FocusZoneDirection.vertical}>
          <DetailsList compact
            layoutMode={DetailsListLayoutMode.justified}
            // selection={selection}
            onItemInvoked={handleItemInvoked}
            selectionMode={SelectionMode.single}
            checkboxVisibility={CheckboxVisibility.hidden}
            enterModalSelectionOnTouch={true}
            selectionPreservedOnEmptyClick={true}
            isHeaderVisible={false}
            columns={[
              {
                key: 'Key', name: 'Key', minWidth: 190, onRender: (item) => {
                  return item
                }
              }]}
            items={expects} />
        </FocusZone>
      </Callout>
    </div>
  </>

}