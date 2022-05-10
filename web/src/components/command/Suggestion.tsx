import { BasePickerListBelow, IBasePickerProps, IBasePickerSuggestionsProps, IInputProps } from '@fluentui/react';
import React from 'react';

export interface ISuggestionProps {

}

export interface ICommandPickerProps extends IBasePickerProps<string> {}

class CommandPicker extends BasePickerListBelow<string, ICommandPickerProps> {}

export const Suggestion = (props: ISuggestionProps) => {

  const onFilterChanged = (filter: string, items: any): Array<string> => {
    console.log(filter, items);
    
    return ['1', '2', '3']
  }

  const pickerSuggestionsProps: IBasePickerSuggestionsProps = {
    suggestionsHeaderText: 'Suggested commands',
    noResultsFoundText: 'No commands found',
  };

  const inputProps: IInputProps = {
    onFocus: () => console.log('onFocus called'),
    onBlur: () => console.log('onBlur called'),
    'aria-label': 'Command picker',
  };

  return <CommandPicker
    removeButtonAriaLabel="Remove"
    // onRenderSuggestionsItem={SuggestedBigItem as any}
    onResolveSuggestions={onFilterChanged}
    // onRenderItem={SelectedDocumentItem}
    // getTextFromItem={getTextFromItem}
    pickerSuggestionsProps={pickerSuggestionsProps}
    selectionAriaLabel="Selected commands"
    selectionRole="group"
    inputProps={inputProps}
  />

}