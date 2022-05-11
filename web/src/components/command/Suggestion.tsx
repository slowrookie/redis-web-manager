import { Callout, CheckboxVisibility, DetailsList, DetailsListLayoutMode, FocusZone, FocusZoneDirection, SelectionMode, TextField } from '@fluentui/react';
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
    [expects, setExpects] = useState<Array<string | number>>([]);

  const handleChange = (e?: ChangeEvent<any>, value?: string) => {
    setCommand(value || "");

    if (value) {
      suggestions(value.trim()).then((v: any) => {
        if (v) {
          const matchToken = value.split(' ')[0].toUpperCase();
          const eps = (v as Array<string>).filter(e => e.indexOf(matchToken) >= 0);
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
        errorMessage={props.errorMessage}
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
        directionalHint={5}
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