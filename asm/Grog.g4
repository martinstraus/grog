grammar Grog;

// Rules
program 
    : instruction+ EOF 
    ; 

instruction
    : increment | decrement
    | arithmeticOperation
    | unaryBooleanOperation | binaryBooleanOperation
    | copyValue | load | store
    | jump
    | input | output
    | stop
    | wait;

load
    : LOAD Register=REGISTER (
        Value=HEXA_BYTE | 
        Address=ABSOLUTE_ADDRESS | 
        Offset=OFFSET_ADDRESS | 
        Pointer=POINTER_ADDRESS
    ) 
    ;

store
    : (Register=REGISTER | Value=HEXA_BYTE) 
      '->' 
      (Address=ABSOLUTE_ADDRESS | Offset=OFFSET_ADDRESS | Pointer=POINTER_ADDRESS) 
    ;

copyValue
    : (SourceRegister=REGISTER '->' DestinationRegister=REGISTER) #CopyRegister
    | (SourceAddress=ABSOLUTE_ADDRESS '->' DestinationAddress=ABSOLUTE_ADDRESS) #CopyAbsoluteToAbsolute
    | (SourceAddress=ABSOLUTE_ADDRESS '->' DestinationOffset=OFFSET_ADDRESS) #CopyAbsoluteToOffset
    | (SourceAddress=ABSOLUTE_ADDRESS '->' DestinationPointer=POINTER_ADDRESS) #CopyAbsoluteToPointer
    | (SourceOffset=OFFSET_ADDRESS '->' DestinationAddress=ABSOLUTE_ADDRESS) #CopyOffsetToAbsolute
    | (SourceOffset=OFFSET_ADDRESS '->' DestinationOffset=OFFSET_ADDRESS) #CopyOffsetToOffset
    | (SourceOffset=OFFSET_ADDRESS '->' DestinationPointer=POINTER_ADDRESS) #CopyOffsetToPointer
    | (SourcePointer=POINTER_ADDRESS '->' DestinationAddress=ABSOLUTE_ADDRESS) #CopyPointerToAbsolute
    | (SourcePointer=POINTER_ADDRESS '->' DestinationOffset=OFFSET_ADDRESS) #CopyPointerToOffset
    | (SourcePointer=POINTER_ADDRESS '->' DestinationPointer=POINTER_ADDRESS) #CopyPointerToPointer
    ;

copyRightToLeft
    : (SourceAddress=ABSOLUTE_ADDRESS | SourceOffset=OFFSET_ADDRESS | SourcePointer=POINTER_ADDRESS) 
      '<->'
      (SourceAddress=ABSOLUTE_ADDRESS | SourceOffset=OFFSET_ADDRESS | SourcePointer=POINTER_ADDRESS) 
    ;

increment
    : INCREMENT Register=REGISTER;

decrement
    : DECREMENT Register=REGISTER;

arithmeticOperation
    : Destination=REGISTER '<-' Left=REGISTER 
      (Operator=ADD| Operator=SUBTRACT| Operator=MULTIPLY | Operator=DIVIDE)
      Right=REGISTER
        ;

unaryBooleanOperation
    : Destination=REGISTER '<-' NOT Operand=REGISTER 
    ;

binaryBooleanOperation
    : Destination=REGISTER '<-' Left=REGISTER 
      (Operator=AND| Operator=OR| Operator=XOR)
      Right=REGISTER
    ;

jump
    : (
        IF Left=REGISTER 
        (
            Operator=EQUAL 
            | Operator=NOT_EQUAL 
            | Operator=GREATER 
            | Operator=GREATER_OR_EQUAL 
            | Operator=LESS
            | Operator=LESS_OR_EQUAL
        ) 
        Right=REGISTER
      )?
      JUMP 
      (Address=ABSOLUTE_ADDRESS | Offset=OFFSET_ADDRESS | Pointer=POINTER_ADDRESS);

input: Destination=REGISTER '<-' Source=DEVICE;

output: Source=REGISTER '->' Destination=DEVICE;

stop
    : STOP
    ;

wait
    : WAIT
    ;

WHITESPACE: [ \r\n\t]+ -> skip;

// Tokens
WS:  [ \t\r\n\u000C]+ -> skip;
COMMENT:   '/*' .*? '*/' -> skip;
LINE_COMMENT :   '//' ~[\r\n]* -> skip;
LOAD: 'load';
STORE: '->';
INCREMENT: 'increment';
DECREMENT: 'decrement';
ADD: '+';
SUBTRACT: '-';
DIVIDE: '/';
MULTIPLY: '*';
EQUAL: '=';
GREATER: '>';
GREATER_OR_EQUAL: '>=';
LESS: '<';
LESS_OR_EQUAL: '<=';
NOT_EQUAL: '!=';
NOT: 'NOT';
AND: 'AND';
XOR: 'XOR';
OR: 'OR';
STOP: 'STOP';
WAIT: 'WAIT';
JUMP: 'JUMP';
IF: 'IF';
HEX_DIGIT: [0-9a-fA-F];
HEXA_BYTE: HEX_DIGIT HEX_DIGIT; 
WORD: HEXA_BYTE HEXA_BYTE; 
REGISTER: 'R' HEX_DIGIT;
DEVICE: 'D' HEXA_BYTE;
ABSOLUTE_ADDRESS: '@'WORD;
OFFSET_ADDRESS: '#'WORD;
POINTER_ADDRESS: '*'WORD;

