<xsl:stylesheet xmlns:xsl="http://www.w3.org/1999/XSL/Transform" version="1.0"
                xmlns:h="http://www.w3.org/1999/xhtml"
                xmlns="http://www.w3.org/1999/xhtml"
                xmlns:sfa="http://developer.apple.com/namespaces/sfa" 
                xmlns:sf="http://developer.apple.com/namespaces/sf" 
                xmlns:sl="http://developer.apple.com/namespaces/sl">

  
  <xsl:template match="/">
    <xsl:apply-templates select="/sl:document/sf:text-storage"/>
  </xsl:template>
  
  <xsl:output method="html"/>
  
  <!-- 
    The main document, we render the styles to CSS classes and they recurse through text-body/section to render the HTML.
    We don't handle list items at this time.
    -->
  <xsl:template match="sf:text-storage">
    <html>
      <head>
        <style>
          <xsl:text>&#10;</xsl:text>
          <xsl:variable name="idref" select="sf:stylesheet-ref/@sfa:IDREF"/>
          <xsl:for-each select="//*[@sfa:ID=$idref]">
            <xsl:apply-templates select=".//sf:paragraphstyle|.//sf:characterstyle"/>
          </xsl:for-each>
        </style>
      </head>
      <body>
        <xsl:apply-templates select="sf:text-body/sf:section"/>
      </body>
    </html>
  </xsl:template>
  
  <xsl:template match="sf:paragraphstyle">
    <xsl:variable name="id" select="@sfa:ID"/>
    <!-- only bother with styles that we use -->
    <xsl:if test="//sf:text-storage//sf:p[@sf:style=$id]">
      <xsl:value-of select="concat('.',@sfa:ID,' {&#10;')"/>
      <xsl:apply-templates select="sf:property-map/*[sf:string|sf:number]"/>
      <xsl:text>}&#10;</xsl:text>
    </xsl:if>
  </xsl:template>
  
  <xsl:template match="sf:characterstyle">
    <xsl:variable name="id" select="@sfa:ID"/>
    <!-- only bother with styles that we use -->
    <xsl:if test="//sf:text-storage//sf:span[@sf:style=$id]">
      <xsl:value-of select="concat('.',@sfa:ID,' {&#10;')"/>
      <xsl:apply-templates select="sf:property-map/*[sf:string|sf:number]"/>
      <xsl:text>}&#10;</xsl:text>
    </xsl:if>
  </xsl:template>
  
  <!-- *** CSS tags -->
  
  <xsl:template match="sf:bold">
    <xsl:text>  font-weight: bold;&#10;</xsl:text>
  </xsl:template>
  
  <xsl:template match="sf:italic">
    <xsl:text>  font-style: italic;&#10;</xsl:text>
  </xsl:template>
  
  <xsl:template match="sf:firstTopicNumber"></xsl:template>
  
  <!-- nobody supports this, but I'll keep it anyway -->
  <xsl:template match="sf:keepWithNext">
    <xsl:text>  page-break-after: avoid;&#10;</xsl:text>
  </xsl:template>
  
  <xsl:template match="sf:keepLinesTogether">
    <xsl:text>  page-break-inside: avoid;&#10;</xsl:text>
  </xsl:template>
  
  <xsl:template match="sf:underline">
    <xsl:text>  text-decoration: underline;&#10;</xsl:text>
  </xsl:template>
  
  <xsl:template match="sf:fontSize">
    <xsl:value-of select="concat('  font-size: ',sf:number/@sfa:number,'px;&#10;')"/>
  </xsl:template>
  
  <xsl:template match="sf:firstLineIndent">
    <xsl:value-of select="concat('  text-indent: ',sf:number/@sfa:number,'px;&#10;')"/>
  </xsl:template>
  <xsl:template match="sf:rightIndent">
    <xsl:value-of select="concat('  margin-right: ',sf:number/@sfa:number,'px;&#10;')"/>
  </xsl:template>
  <xsl:template match="sf:leftIndent">
    <xsl:value-of select="concat('  margin-left: ',sf:number/@sfa:number,'px;&#10;')"/>
  </xsl:template>
  
  <xsl:template match="sf:spaceAfter">
    <xsl:value-of select="concat('  margin-bottom: ',sf:number/@sfa:number,'px;&#10;')"/>
  </xsl:template>
  
  <xsl:template match="sf:fontName">
    <xsl:value-of select="concat('  font-family: ',sf:string/@sfa:string,';&#10;')"/>
  </xsl:template>
  
  <xsl:template match="sf:showInTOC">
    <xsl:variable name="value" select="number(sf:number/@sfa:number)"/>
    <xsl:value-of select="concat('/* this is a H',$value,' heading */&#10;')"/>
  </xsl:template>
  
  <xsl:template match="sf:section">
    <section title="{@sf:name}" class="{@sf:style}">
      <xsl:apply-templates select="sf:layout/*"/>
    </section>
  </xsl:template>
  
  <xsl:template match="sf:alignment">
    <xsl:variable name="value" select="number(sf:number/@sfa:number)"/>
    <xsl:choose>
      <xsl:when test="$value=2"> 
        <xsl:text>  text-align: centered;&#10;</xsl:text>
      </xsl:when>
      <xsl:when test="$value=3"> 
        <xsl:text>  text-align: justify;&#10;</xsl:text>
      </xsl:when>
      <xsl:otherwise>
        <xsl:message>Unhandled alignment value: <xsl:value-of select="$value"/></xsl:message>
      </xsl:otherwise>
    </xsl:choose>
  </xsl:template>
  
  <xsl:template match="*">
    <xsl:message>Unhandled tag: <xsl:value-of select="name()"/>&#10;</xsl:message>
  </xsl:template>
  
  <!-- *** Document content tags -->
  
  <xsl:template match="sf:p">
    <p><xsl:call-template name="recurseWithStyle"/></p>
  </xsl:template>

<!-- This doesn't work because there is a p for each li and no way to collect them together.
  
  <xsl:template match="sf:p[.//sf:tab]">
    <xsl:variable name="label" select="sf:span/sf:tab[1]/following-sibling::text()"/>
    <xsl:variable name="num" select="number($label)"/>
    <xsl:choose>
      <xsl:when test="$num > 0">
        HANDLE OL
      </xsl:when>
      <xsl:otherwise>
        HANLDLE UL
      </xsl:otherwise>
    </xsl:choose>
  </xsl:template>
-->
  
  <xsl:template match="sf:tab">
    <!-- FIXME - these are OL/UL lists, we should special case paragraphs containing sf:tab and drop this rule. -->
    <xsl:text>&#9;</xsl:text>
  </xsl:template>
  
  <xsl:template match="sf:page-start"><a id="page{@sf:page-index}"></a></xsl:template>
  <xsl:template name="recurseWithStyle">
    <xsl:if test="@sf:style">
      <!-- FIXME - transform style name? -->
      <xsl:attribute name="class"><xsl:value-of select="@sf:style"/></xsl:attribute>
    </xsl:if>
    <xsl:apply-templates/>
  </xsl:template>
  
  <xsl:template match="sf:span">
    <span><xsl:call-template name="recurseWithStyle"/></span>
  </xsl:template>
  
  <xsl:template match="sf:link">
    <a href="{@href}"><xsl:call-template name="recurseWithStyle"/></a>
  </xsl:template>
  
  <xsl:template match="sf:link-ref">
    <xsl:variable name='idref' select='@sfa:IDREF'/>
    <xsl:variable name='href' select='//sf:link[@sfa:ID=$idref]/@href'/>
    <a href="{$href}"><xsl:call-template name="recurseWithStyle"/></a>
  </xsl:template>
  
  <xsl:template match="sf:attachment-ref">
    <xsl:variable name="idref" select="@sfa:IDREF"/>
    <xsl:apply-templates select="//sf:attachment[@sfa:ID=$idref]"/>
  </xsl:template>
  
  
  <xsl:template match="sf:attachment[sf:media]">
    <!-- FIXME size -->
    <xsl:apply-templates select=".//sf:unfiltered-ref"/>
    <xsl:apply-templates select=".//sf:unfiltered"/>
  </xsl:template>
  
  <xsl:template match="sf:attachment[sf:tabular-info]">
    <xsl:variable name="ref" select="sf:tabular-info/sf:tabular-model-ref/@sfa:IDREF"/>
    <xsl:apply-templates select="//sf:tabular-model[@sfa:ID=$ref]"/>
  </xsl:template>
  
  <xsl:template match="sf:tabular-model">
    <xsl:variable name="rows" select="sf:grid/sf:rows/*"/>
    
    <xsl:variable name="cols" select="sf:grid/sf:columns/*"/>

    <table>
      <xsl:for-each select="$rows">
        <tr>
          <xsl:variable name="row" select="position()"/>
          <xsl:for-each select="$cols">
            <td>
              <xsl:variable name="col" select="position()"/>
              <xsl:variable name="cell" select="ancestor::sf:grid/sf:datasource/*[@sf:row = $row - 1][@sf:col = $col - 1]"/>
              <xsl:apply-templates select="$cell//sf:text-body/*"/>
            </td>
          </xsl:for-each>
        </tr>
      </xsl:for-each>
    </table>
    
  </xsl:template>
  
  <xsl:template match="sf:attachment">
    <xsl:message>Unhandled attachment <xsl:value-of select="@sfa:ID"/></xsl:message>
  </xsl:template>
  
  <xsl:template match="sf:unfiltered-ref">
    <xsl:variable name="ref" select="sfa:IDREF"/>
    <xsl:apply-templates select="//sf:unfiltered[sfa:ID=$ref]"/>
  </xsl:template>
  
  <xsl:template match="sf:unfiltered">
      <!-- this is naive - sometimes there is an unfiltered image ref to a previous image if the image is a dup -->
      <xsl:variable name="path" select="sf:data/@sf:path"/>
      <xsl:if test="not($path)">
        <xsl:message>Missing path for image in attachment <xsl:value-of select="ancestor::sf:attachment/@sfa:ID"/></xsl:message>
      </xsl:if>
      <img src="{$path}" 
           width="{sf:size/@sfa:w}" 
           height="{sf:size/@sfa:h}"/>
  </xsl:template>
  
  <xsl:template match="sf:lnbr">
    <br/>
  </xsl:template>
  
  <!-- ignored tags -->
  <xsl:template match="sf:br|sf:selection-start|sf:selection-end|sf:container-hint|sf:insertion-point|sf:crbr"/>
  
  <!-- Debugging helpers -->
  
  <xsl:template mode="list" match="*">
    <xsl:text>&#10;</xsl:text>
    <xsl:for-each select="*">
      <xsl:value-of select="name()"/><xsl:text>&#10;</xsl:text>
    </xsl:for-each>
  </xsl:template>
  
  <xsl:template name="dump">
    <xsl:for-each select="*">
      <xsl:apply-templates select="." mode="copy"/>
      <xsl:text>&#10;&#10;</xsl:text>
    </xsl:for-each>
  </xsl:template>
  
  <xsl:template mode="copy" match="*|@*">
    <xsl:copy>
      <xsl:apply-templates select="node()|@*" mode="copy"/>
    </xsl:copy>
  </xsl:template>
  
</xsl:stylesheet>
